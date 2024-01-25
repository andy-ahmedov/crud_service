package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	domain "github.com/andy-ahmedov/crud_service/internal/Domain"
	"github.com/andy-ahmedov/crud_service/internal/Repository/psql"
	"github.com/andy-ahmedov/crud_service/pkg/postgres"
)

func main() {
	conn, err := postgres.ConnectToDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTED")

	db := psql.NewBook(conn)

	book := domain.Book{
		Title:       "Something",
		Author:      "Dan Brown",
		PublishDate: time.Time{},
		Rating:      5}

	err = db.Create(context.Background(), &book)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(book.ID)

	// var id int64

	// id = 3

	bookOne, err := db.GetByID(context.TODO(), 1)
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

	updBook := domain.UpdateBookInput{
		Title:       &title,
		Author:      &author,
		PublishDate: &time.Time{},
		Rating:      &rating,
	}

	err = db.Update(context.TODO(), 7, updBook)
	if err != nil {
		log.Fatal(err)
	}

	books, err := db.GetAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)

}
