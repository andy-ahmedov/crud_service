package service

import (
	"context"
	"database/sql"
	"errors"
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
	// хешируем пароль
	// отправляем данные для входа наверх для получения юзера
	// если получаем ошибку ErrNoRows, то возвращаем ошибку отсутствия пользователя с этими данными
	// генерим токен используя метод jwt.NewWithClaims
	// подписываем с помощью секрета получившийся токен используя метод SignedString
	// и возвращаем его
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
		Id:        strconv.Itoa(int(getIt.ID)),
	})

	return token.SignedString(u.hmacSecret)
}
