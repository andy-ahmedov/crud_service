package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/andy-ahmedov/crud_service/internal/service"
	"github.com/gin-gonic/gin"
)

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

	router.Run()

	return router
}

func (h Handler) createBook(c *gin.Context) {
	var book domain.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.booksService.Create(context.TODO(), &book)
	if err != nil {
		log.Println("ginCreateBook() error: ", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s\n", err))
		return
	}

	c.String(http.StatusOK, "The data has been successfully written.\n")
}

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.booksService.GetAll(c)
	if err != nil {
		log.Println("ginGetAllBooks() error:", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s\n", err))
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) getBookByID(c *gin.Context) {
	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s\n", err))
		return
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			c.String(http.StatusBadRequest, fmt.Sprintf("error: %s\n", err))
			return
		}
		log.Println("getByID() error:", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s\n", err))
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) deleteBook(c *gin.Context) {
	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s\n", err))
		return
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s\n", err))
		return
	}

	c.String(http.StatusOK, "The row with the given ID was successfully deleted.\n")
}

func (h *Handler) updateBook(c *gin.Context) {
	var updBook domain.UpdateBookInput

	if err := c.ShouldBindJSON(&updBook); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = h.booksService.Update(context.TODO(), id, updBook)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Oups(( error: %s\n", err))
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
