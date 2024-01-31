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

func (u *UserRepository) GetByCredential(ctx context.Context, email string, password string) (domain.User, error) {
	var user domain.User
	request := `SELECT id, name, email, password, registered_at FROM users WHERE email=$1 AND password=$2`
	err := u.db.QueryRow(ctx, request, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAt)

	return user, err
}

// GetByCredential (domain.User, error)
// отправляем запрос с эмейлом и паролем в базу
// Получаем обратно структуру User со всеми полями
