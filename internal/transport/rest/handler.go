package rest

import (
	"context"
	"errors"
	"strconv"

	_ "github.com/andy-ahmedov/crud_service/docs"
	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type BooksRepository interface {
	Create(ctx context.Context, book *domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, updBook domain.UpdateBookInput) error
}

type UserRepository interface {
	SignUp(ctx context.Context, inp domain.SignUpInput) error
	SignIn(ctx context.Context, inp domain.SignInInput) (string, error)
}

type errResponse struct {
	Message string
}

type Handler struct {
	booksService BooksRepository
	userService  UserRepository
}

func NewHandler(books BooksRepository, users UserRepository) *Handler {
	return &Handler{
		booksService: books,
		userService:  users,
	}
}

func (h Handler) InitGinRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/books", h.createBook)
	router.GET("/books", h.getAllBooks)
	router.GET("/books/:id", h.getBook)
	router.DELETE("/books/:id", h.deleteBook)
	router.PUT("books/:id", h.updateBook)

	router.POST("/auth/sign-up", h.signUp)
	router.POST("/auth/sign-in", h.signIn)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

	return router
}

func getIDFromRequest(param string) (int64, error) {
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("ID can't be 0")
	}

	return id, nil
}
