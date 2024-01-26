package service

import (
	"context"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
)

type BooksInterface interface {
	Create(ctx context.Context, book *domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, updBook domain.UpdateBookInput) error
}

type BookStorage struct {
	repo BooksInterface
}

func NewBooksStorage(repo BooksInterface) *BookStorage {
	return &BookStorage{repo: repo}
}

func (b BookStorage) Create(ctx context.Context, book *domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return b.repo.Create(ctx, book)
}

func (b *BookStorage) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *BookStorage) GetAll(ctx context.Context) ([]domain.Book, error) {
	return b.repo.GetAll(ctx)
}

func (b *BookStorage) Delete(ctx context.Context, id int64) error {
	return b.repo.Delete(ctx, id)
}

func (b *BookStorage) Update(ctx context.Context, id int64, updBook domain.UpdateBookInput) error {
	return b.repo.Update(ctx, id, updBook)
}
