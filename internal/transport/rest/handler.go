package rest

import (
	"context"
	"errors"
	"strconv"

	_ "github.com/andy-ahmedov/crud_service/docs"
	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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
	SignIn(ctx context.Context, inp domain.SignInInput) (domain.User, error)
}

type errResponse struct {
	Message string
}

type Handler struct {
	booksService BooksRepository
	userService  UserRepository
	store        *sessions.CookieStore
}

func NewHandler(books BooksRepository, users UserRepository, store *sessions.CookieStore) *Handler {
	return &Handler{
		booksService: books,
		userService:  users,
		store:        store,
	}
}

func (h Handler) InitGinRouter() *gin.Engine {
	router := gin.Default()

	router.Use(loggingMiddleware)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/logout", h.logout)
	}

	books := router.Group("/books", h.authMiddleware)
	{
		books.POST("", h.createBook)
		books.GET("", h.getAllBooks)

		id := books.Group("/:id")
		{
			id.GET("", h.getBook)
			id.DELETE("", h.deleteBook)
			id.PUT("", h.updateBook)
		}
	}

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
