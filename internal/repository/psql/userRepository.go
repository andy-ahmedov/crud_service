package psql

import (
	"context"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	_ "github.com/andy-ahmedov/crud_service/internal/transport/rest"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user domain.User) error {
	request := `INSERT INTO users(name, email, password, registered_at) VALUES($1, $2, $3, $4) RETURNING id`
	err := u.db.QueryRow(ctx, request, user.Name, user.Email, user.Password, user.RegisteredAt).Scan(&user.ID)

	return err
}
