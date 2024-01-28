package main

import (
	"context"
	"log"
	"net/http"
	"time"

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

func main() {
	db, err := postgres.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	booksRepo := psql.NewBook(db)
	booksService := service.NewBooksStorage(booksRepo)
	handler := rest.NewHandler(booksService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitGinRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
