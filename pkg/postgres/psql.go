package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
)

type ConnectionInfo struct {
	Port     int
	Host     string
	UserName string
	DBname   string
	SSLMode  string
	Password string
}

func ConnectToDB() (*pgx.Conn, error) {
	urlExample := "postgres://postgres:mark@localhost:5434/booking"
	// conStr := fmt.Sprintf("port=%s host=%s user=%s dbname=%s sslmode%s password=%s", "5434", "localhost", "postgres", "booking", "disable", "mark")
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

func GetConfig() (ConnectionInfo, error) {
	var config ConnectionInfo

	file, err := os.ReadFile("/home/andy/github.com/andy-ahmedov/crud_service/config.yaml")
	if err != nil {
		return config, err
	}

	yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
