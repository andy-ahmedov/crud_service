package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary CreateBook
// @Security ApiKeyAuth
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
		logError("createBook", "writing data to a structure", err)
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err := h.booksService.Create(context.TODO(), &book)
	if err != nil {
		logError("createBook", "service error", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": book.ID})
}

// @Summary getAllBooks
// @Security ApiKeyAuth
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
		logError("getAllBooks", "reading data from a database", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary GetBookByID
// @Security ApiKeyAuth
// @Tags id
// @Description Retrieves a book by ID. If the book is not found, returns an error.
// @ID get-book-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} domain.Book "OK"
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /books/{id} [get]
func (h *Handler) getBook(c *gin.Context) {
	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		logError("getBook", "reading id from request", err)
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			logError("getBook", "there is no book with the given identifier", err)
			c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
			return
		}
		logError("getBook", "reading data from a database", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary deleteBook
// @Security ApiKeyAuth
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
		logError("deleteBook", "reading id from request", err)
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		logError("deleteBook", "deleting data from the database", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.String(http.StatusOK, "The row with the given ID was successfully deleted.\n")
}

// @Summary updateBook
// @Security ApiKeyAuth
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
		logError("updateBook", "writing data to a structure", err)
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		logError("updateBook", "reading id from request", err)
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err = h.booksService.Update(context.TODO(), id, updBook)
	if err != nil {
		logError("updateBook", "service error", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.String(http.StatusOK, "The book with the given ID has been successfully updated.\n")
}
