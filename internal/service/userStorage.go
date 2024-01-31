package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/golang-jwt/jwt"
)

type Users struct {
	repo   UserStorage
	hasher PasswordHasher

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

func NewUsers(repo UserStorage, hasher PasswordHasher, secret []byte, ttl time.Duration) *Users {
	return &Users{
		repo:       repo,
		hasher:     hasher,
		hmacSecret: secret,
		tokenTtl:   ttl,
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

func (u *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, error) {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return "", err
	}

	getIt, err := u.repo.GetByCredential(ctx, inp.Email, password)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", domain.ErrUserNotFound
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(u.tokenTtl).Unix(),
		Subject:   strconv.Itoa(int(getIt.ID)),
	})

	return token.SignedString(u.hmacSecret)
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

// ParseToken(ctx context.Context, token string) (int64, error)
// вызываем функцию jwt.Parse(..., func(token *jwt.Token) (interface{}, error) {
// найти в документации jwt (example parsing and validating token HMAC)
// в конце анонимной функции возвращаем наш секрет
// Проверка на ошибку
// Проверка на валидность
// Достаем клеймсы из t.Claims(jwt.MapClaims). При ошибке "invalid claims"
// Достаем сабджект из клеймса выше используя ["sub"].(string). При ошибке "invalid subject"
// Преобразовываем полученный сабджект в int. При ошибке "invalid subject"
// Возвращаем int64(id)
