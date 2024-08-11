package rest

import (
	"net/http"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
)

type Handler struct {
	booksHandler *BooksHandler
	usersHandler *UsersHandler
}

func NewHandler(booksService *service.Books, usersService *service.Users) *Handler {
	return &Handler{
		booksHandler: NewBooksHandler(booksService),
		usersHandler: NewUsersHandler(usersService),
	}
}

func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Books routes
	// mux.HandleFunc("/books", h.methodHandler(h.booksHandler.CreateBook, http.MethodPost))
	// mux.HandleFunc("/books/", h.handleBookID(h.booksHandler.GetBookByID, http.MethodGet))
	// mux.HandleFunc("/books", h.methodHandler(h.booksHandler.GetAllBooks, http.MethodGet))
	// mux.HandleFunc("/books/", h.handleBookID(h.booksHandler.UpdateBook, http.MethodPut))
	// mux.HandleFunc("/books/", h.handleBookID(h.booksHandler.DeleteBook, http.MethodDelete))

	// Users routes
	mux.HandleFunc("/signup", h.methodHandler(h.usersHandler.SignUp, http.MethodPost))
	mux.HandleFunc("/signin", h.methodHandler(h.usersHandler.SignIn, http.MethodPost))

	return mux
}

// methodHandler is a helper function to restrict handler to a specific HTTP method.
func (h *Handler) methodHandler(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}

// handleBookID extracts the book ID from the URL and calls the handler.
func (h *Handler) handleBookID(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// Extract the book ID from the URL
		id := r.URL.Path[len("/books/"):]
		if id == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// Call the original handler with ID as URL param
		handlerFunc(w, r)
	}
}
