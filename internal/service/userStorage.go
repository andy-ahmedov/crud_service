package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
)

// добавляем новое поле sessionRepo SessionRepository
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

type SessionRepository interface {
	Create(ctx context.Context, token domain.RefreshSession) error
	Get(ctx context.Context, token string) (domain.RefreshSession, error)
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

func (u *Users) SignIn(ctx context.Context, inp domain.SignInInput) (domain.User, error) {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return domain.User{}, err
	}

	getIt, err := u.repo.GetByCredential(ctx, inp.Email, password)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return getIt, err
}
