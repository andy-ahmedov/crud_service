package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	audit "github.com/andy-ahmedov/audit_log_server/pkg/domain"
	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=userStorage.go -destination=mocks/mock.go

type AuditClient interface {
	SendLogRequest(ctx context.Context, req audit.LogItem) error
}

type Users struct {
	Repo        UserStorage
	Hasher      PasswordHasher
	SessionRepo SessionRepository
	AuditClient AuditClient

	HmacSecret []byte
	TokenTtl   time.Duration
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type UserStorage interface {
	CreateUser(ctx context.Context, inp domain.User) error
	GetByCredential(ctx context.Context, email string, passwords string) (domain.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, token domain.RefreshSession) error
	Get(ctx context.Context, token string) (domain.RefreshSession, error)
}

// также добавляем новое поле в NewUsers
func NewUsers(repo UserStorage, hasher PasswordHasher, sessionRepo SessionRepository, auditClient AuditClient, secret []byte, ttl time.Duration) *Users {
	return &Users{
		Repo:        repo,
		Hasher:      hasher,
		HmacSecret:  secret,
		TokenTtl:    ttl,
		SessionRepo: sessionRepo,
		AuditClient: auditClient,
	}

}

func (u *Users) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password, err := u.Hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	if err := u.Repo.CreateUser(ctx, user); err != nil {
		return err
	}

	user, err = u.Repo.GetByCredential(ctx, inp.Email, password)
	if err != nil {
		return err
	}

	if err := u.AuditClient.SendLogRequest(ctx, audit.LogItem{
		Action:    audit.ACTION_REGISTER,
		Entity:    audit.ENTITY_USER,
		EntityID:  user.ID,
		Timestamp: time.Now(),
	}); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "User.SignUp",
		}).Error("failed to send log request:", err)
	}
	return err
}

func (u *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error) {
	password, err := u.Hasher.Hash(inp.Password)
	if err != nil {
		return "", "", err
	}

	getIt, err := u.Repo.GetByCredential(ctx, inp.Email, password)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", "", domain.ErrUserNotFound
		}
		return "", "", err
	}

	return u.generateTokens(ctx, getIt.ID)
}

func (u *Users) ParseToken(ctx context.Context, token string) (int64, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return u.HmacSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject := claims["sub"].(string)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return int64(id), nil
}

func (u *Users) generateTokens(ctx context.Context, userID int64) (string, string, error) {
	notSigned := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(u.TokenTtl).Unix(),
		Subject:   strconv.Itoa(int(userID)),
	})

	accessToken, err := notSigned.SignedString(u.HmacSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	session := domain.RefreshSession{
		UserID:    userID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}

	err = u.SessionRepo.Create(ctx, session)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func newRefreshToken() (string, error) {
	refresh := make([]byte, 32)

	newSourse := rand.NewSource(time.Now().Unix())
	r := rand.New(newSourse)

	_, err := r.Read(refresh)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", refresh), err
}

func (u *Users) RefreshTokens(ctx context.Context, refreshToken string) (string, string, error) {
	session, err := u.SessionRepo.Get(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}

	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", domain.ErrRefreshTokenExpired
	}

	return u.generateTokens(ctx, session.UserID)
}
