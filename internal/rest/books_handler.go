package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
)

type BooksHandler struct {
	service *service.Books
}

func NewBooksHandler(service *service.Books) *BooksHandler {
	return &BooksHandler{service: service}
}

func (h *BooksHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(r.Context(), book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BooksHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	book, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (h *BooksHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (h *BooksHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(r.Context(), id, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BooksHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	if err := h.service.Delete(r.Context(), id); err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
