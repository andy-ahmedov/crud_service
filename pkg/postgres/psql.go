package postgres

import (
	"context"
	"fmt"

	"github.com/andy-ahmedov/crud_service/internal/config"
	"github.com/jackc/pgx/v5"
	_ "gopkg.in/yaml.v3"
)

// urlExample := "postgres://postgres:mark@localhost:5434/booking"

func ConnectToDB(cfg config.Postgres) (*pgx.Conn, error) {
	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.TODO())
	if err != nil {
		return nil, err
	}
	fmt.Println("CONNECTED")

	return conn, err
}
