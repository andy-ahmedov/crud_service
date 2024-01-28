package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	_ "github.com/andy-ahmedov/crud_service/docs"
	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/andy-ahmedov/crud_service/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type errResponse struct {
	Message string
}

type Handler struct {
	booksService service.BooksInterface
}

func NewHandler(books service.BooksInterface) *Handler {
	return &Handler{
		booksService: books,
	}
}

func (h Handler) InitGinRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/books", h.createBook)
	router.GET("/books", h.getAllBooks)
	router.GET("/books/:id", h.getBookByID)
	router.DELETE("/books/:id", h.deleteBook)
	router.PUT("books/:id", h.updateBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

	return router
}

// @Summary CreateBook
// @Tags books
// @Description Adding a book to the database.
// @ID add-book
// @Accept json
// @Produce json
// @Param input body domain.Book true "Book information"
// @Success 200 {string} gin.H "The data has been successfully written."
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /books [post]
func (h Handler) createBook(c *gin.Context) {
	var book domain.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err := h.booksService.Create(context.TODO(), &book)
	if err != nil {
		// log.Println("ginCreateBook() error: ", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": book.ID})
}

// @Summary getAllBooks
// @Tags books
// @Description Getting all books.
// @ID get-all-books
// @Produce json
// @Success 200 {array} domain.Book "Books have been successfully received."
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /books [get]
func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.booksService.GetAll(c)
	if err != nil {
		// log.Println("ginGetAllBooks() error:", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary GetBookByID
// @Description Retrieves a book by ID. If the book is not found, returns an error.
// @ID get-book-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} domain.Book "OK"
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /books/{id} [get]
func (h *Handler) getBookByID(c *gin.Context) {
	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
			return
		}
		// log.Println("getByID() error:", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary deleteBook
// @Tags id
// @Description Deleting a book by ID.
// @ID delete-book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {string} string "The data has been successfully written."
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /books/{id} [delete]
func (h *Handler) deleteBook(c *gin.Context) {
	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.String(http.StatusOK, "The row with the given ID was successfully deleted.\n")
}

// @Summary updateBook
// @Tags id
// @Description Updating book data by ID.
// @ID update-book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param updateBook body domain.UpdateBookInput true "Book update information"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /books/{id} [put]
func (h *Handler) updateBook(c *gin.Context) {
	var updBook domain.UpdateBookInput

	if err := c.ShouldBindJSON(&updBook); err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err = h.booksService.Update(context.TODO(), id, updBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.String(http.StatusOK, "The book with the given ID has been successfully updated.\n")
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
