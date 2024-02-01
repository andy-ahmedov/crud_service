package psql

import (
	"context"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/jackc/pgx/v5"
)

// добавить структуру Tokens содержащую в себе поле db
//
// NewTokens
//

// метод Create(ctx, token domain.RefreshSession) error
// запрос Exec в бд (INSERT INTO (юзерАйди, токен и время окончания) VALUES())
// возвращаем ошибку

// метод Get(ctx, token string) (domainRefreshSession) error {
// создаем новую переменную типа domain RefreshSession
// запрос QueryRow в бд (SELECT айди, юзерАйди токен и время окончания FROM refresh_tokens WHERE токен равен $1) c последующим сканированием всех полей. Проверка на ошибки
// удаляем строку из таблицы refreshTokens где user_id равно возвращенной userID используя Exec и  DELETE FROM
// возвращаем созданную переменную и ошибку

type Tokens struct {
	db *pgx.Conn
}

func NewTokens(db *pgx.Conn) *Tokens {
	return &Tokens{db: db}
}

func (t *Tokens) Create(ctx context.Context, token domain.RefreshSession) error {
	_, err := t.db.Exec(ctx, "INSERT INTO refresh_tokens(user_id, token, expires_at) VALUES($1, $2, $3)", token.UserID, token.Token, token.ExpiresAt)

	return err
}

func (t *Tokens) Get(ctx context.Context, token string) (domain.RefreshSession, error) {
	var session domain.RefreshSession

	err := t.db.QueryRow(ctx, "SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1", token).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresAt)
	if err != nil {
		return session, err
	}
	_, err = t.db.Exec(ctx, "DELETE FROM refresh_tokens WHERE user_id=$1", session.UserID)

	return session, err
}
