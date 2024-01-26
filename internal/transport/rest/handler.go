package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/andy-ahmedov/crud_service/internal/service"
	"github.com/gorilla/mux"
)

// type BooksInterface interface {
// 	Create(ctx context.Context, book *domain.Book) error
// 	GetByID(ctx context.Context, id int64) (domain.Book, error)
// 	GetAll(ctx context.Context) ([]domain.Book, error)
// 	Delete(ctx context.Context, id int64) error
// 	Update(ctx context.Context, id int64, updBook domain.UpdateBookInput) error
// }

type Handler struct {
	booksService service.BooksInterface
}

func NewHandler(books service.BooksInterface) *Handler {
	return &Handler{
		booksService: books,
	}
}

func (h Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	books := r.PathPrefix("/books").Subrouter()
	{
		books.HandleFunc("", h.createBook).Methods(http.MethodPost)
		books.HandleFunc("", h.getAllBooks).Methods(http.MethodGet)
		books.HandleFunc("/{id: [0-9]+}", h.getBookByID).Methods(http.MethodGet)
		books.HandleFunc("/{id: [0-9]+}", h.deleteBook).Methods(http.MethodDelete)
		books.HandleFunc("/{id: [0-9]+}", h.updateBook).Methods(http.MethodPut)
	}

	return r
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book domain.Book

	err = json.Unmarshal(reqBytes, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Create(context.TODO(), &book)
	if err != nil {
		log.Println("createBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.booksService.GetAll(context.TODO())
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(books)
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("getByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		log.Println("getBookByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		log.Println("deleteBook() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		log.Println("deleteBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("updateBook() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updBook domain.UpdateBookInput

	err = json.Unmarshal(reqBytes, &updBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := getIDFromRequest(r)
	if err != nil {
		log.Println("updateBook() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Update(context.TODO(), id, updBook)
	if err != nil {
		log.Println("updateBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getIDFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("ID can't be 0")
	}

	return id, nil
}
