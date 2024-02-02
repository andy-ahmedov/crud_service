package main

import (
	"context"
	"encoding/gob"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"

	"github.com/andy-ahmedov/crud_service/internal/config"
	"github.com/andy-ahmedov/crud_service/internal/domain"
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
	userRepo := psql.NewUserRepository(db)

	booksService := service.NewBooksStorage(booksRepo)
	userService := service.NewUsers(userRepo, hasher, []byte(cfg.Secret), cfg.TokenTTL)

	sessionStore := sessions.NewCookieStore([]byte(cfg.Secret))
	sessionStore.Options.HttpOnly = true
	gob.Register(&domain.User{})

	handler := rest.NewHandler(booksService, userService, sessionStore)

	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: handler.InitGinRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
