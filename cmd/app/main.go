package main

import (
	"context"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/andy-ahmedov/crud_service/internal/config"
	"github.com/andy-ahmedov/crud_service/internal/repository/psql"
	"github.com/andy-ahmedov/crud_service/internal/service"
	"github.com/andy-ahmedov/crud_service/internal/transport/rest"
	"github.com/andy-ahmedov/crud_service/pkg/postgres"
)

// @title CRUD API Service
// @version 1.2
// @description Service implementing crud operations

// @contact.name Andy Ahmedov
// @contact.url https://github.com/andy-ahmedov
// @contact.email andy.ahmedov@gmail.com

// @host localhost:8080
// @BasePath /

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectToDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	booksRepo := psql.NewBook(db)
	booksService := service.NewBooksStorage(booksRepo)
	handler := rest.NewHandler(booksService)

	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: handler.InitGinRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
