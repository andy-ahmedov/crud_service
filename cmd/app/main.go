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
	"github.com/andy-ahmedov/crud_service/pkg/hash"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
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

	hasher := hash.NewSHA1Hasher(cfg.Salt)

	booksRepo := psql.NewBookRepository(db)
	sessionRepo := psql.NewTokens(db)
	// добавить репозиторий токена. Включить его в параметры NewUsers
	booksService := service.NewBooksStorage(booksRepo)

	userRepo := psql.NewUserRepository(db)
	userService := service.NewUsers(userRepo, hasher, sessionRepo, []byte(cfg.Secret), cfg.TokenTTL)

	handler := rest.NewHandler(booksService, userService)

	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: handler.InitGinRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
