package psql

import (
	"context"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/jackc/pgx/v5"
)

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
