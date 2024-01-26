package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/andy-ahmedov/crud_service/internal/repository/psql"
	"github.com/andy-ahmedov/crud_service/internal/service"
	"github.com/andy-ahmedov/crud_service/internal/transport/rest"
	"github.com/andy-ahmedov/crud_service/pkg/postgres"
)

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
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	tiime := time.Now()
	book := domain.Book{
		Title:       "Something",
		Author:      "Dan Brown",
		PublishDate: tiime,
		Rating:      5,
	}

	err = booksService.Create(context.Background(), &book)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(book.ID)

	bookOne, err := booksService.GetByID(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bookOne)

	// err = db.Delete(context.TODO(), 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	title := "Tommy"
	author := "JimBim"
	rating := 3
	tiime = time.Now()

	updBook := domain.UpdateBookInput{
		Title:       &title,
		Author:      &author,
		PublishDate: &tiime,
		Rating:      &rating,
	}

	err = booksService.Update(context.TODO(), 6, updBook)
	if err != nil {
		log.Fatal(err)
	}

	books, err := booksService.GetAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)

}
