package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/golang-jwt/jwt"
)

// добавляем новое поле sessionRepo SessionRepository
type Users struct {
	repo        UserStorage
	hasher      PasswordHasher
	sessionRepo SessionRepository

	hmacSecret []byte
	tokenTtl   time.Duration
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
func NewUsers(repo UserStorage, hasher PasswordHasher, sessionRepo SessionRepository, secret []byte, ttl time.Duration) *Users {
	return &Users{
		repo:        repo,
		hasher:      hasher,
		hmacSecret:  secret,
		tokenTtl:    ttl,
		sessionRepo: sessionRepo,
	}

}

func (u *Users) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	return u.repo.CreateUser(ctx, user)
}

func (u *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error) {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return "", "", err
	}

	getIt, err := u.repo.GetByCredential(ctx, inp.Email, password)
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
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return u.hmacSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("Invalid token")
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
		ExpiresAt: time.Now().Add(u.tokenTtl).Unix(),
		Subject:   strconv.Itoa(int(userID)),
	})

	accessToken, err := notSigned.SignedString(u.hmacSecret)
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

	err = u.sessionRepo.Create(ctx, session)
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
	session, err := u.sessionRepo.Get(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}

	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", domain.ErrRefreshTokenExpired
	}

	return u.generateTokens(ctx, session.UserID)
}
