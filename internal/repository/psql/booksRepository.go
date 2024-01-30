package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Books struct {
	db *pgx.Conn
}

func NewBookRepository(db *pgx.Conn) *Books {
	return &Books{db}
}

func (b *Books) Create(ctx context.Context, book *domain.Book) error {
	request := `INSERT INTO books(title, author, publish_date, rating) VALUES($1, $2, $3, $4) RETURNING id`
	if err := b.db.QueryRow(ctx, request, book.Title, book.Author, book.PublishDate, book.Rating).Scan(&book.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.WithFields(log.Fields{
				"file":     "booksRepository.go",
				"func":     "Create",
				"problem":  "SQL Request Error",
				"error":    pgErr.Message,
				"detail":   pgErr.Detail,
				"where":    pgErr.Where,
				"code":     pgErr.Code,
				"SQLState": pgErr.SQLState(),
			}).Error(err)
			return nil
		}
		return err
	}

	return nil
}

func (b *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	var book domain.Book
	request := fmt.Sprintf(`SELECT * FROM books WHERE id=$1`)
	err := b.db.QueryRow(ctx, request, id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)
	if err == sql.ErrNoRows {
		return book, domain.ErrBookNotFound
	}
	return book, err
}

func (b *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	books := make([]domain.Book, 0)
	request := "SELECT * FROM books"

	rows, err := b.db.Query(ctx, request)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book domain.Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Books) Delete(ctx context.Context, id int64) error {
	_, err := b.db.Exec(ctx, "DELETE FROM books WHERE id=$1", id)
	return err
}

func (b *Books) Update(ctx context.Context, id int64, upd domain.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if upd.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *upd.Title)
		argId++
	}

	if upd.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *upd.Author)
		argId++
	}

	if upd.PublishDate != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, *upd.PublishDate)
		argId++
	} else {
		t := time.Now()
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, &t)
		argId++
	}

	if upd.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *upd.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE books SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := b.db.Exec(ctx, query, args...)
	return err
}
