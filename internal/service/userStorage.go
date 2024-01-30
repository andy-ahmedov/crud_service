package service

import (
	"context"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
)

type Users struct {
	repo   UserStorage
	hasher PasswordHasher

	hmacSecret []byte
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type UserStorage interface {
	CreateUser(ctx context.Context, inp domain.User) error
}

func NewUsers(repo UserStorage, hasher PasswordHasher, secret []byte) *Users {
	return &Users{
		repo:       repo,
		hasher:     hasher,
		hmacSecret: secret,
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
