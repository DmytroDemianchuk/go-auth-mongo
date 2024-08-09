package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
)

// Handler manages HTTP requests.
type Handler struct {
	booksService *service.BooksService
}

// NewHandler creates a new Handler.
func NewHandler(booksService *service.BooksService) *Handler {
	return &Handler{booksService: booksService}
}

// InitRouter initializes the HTTP router.
func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", h.getBooks)
	mux.HandleFunc("/books/create", h.createBook)
	return mux
}

func (h *Handler) getBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	books, err := h.booksService.ListBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book interface{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.booksService.AddBook(r.Context(), book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
