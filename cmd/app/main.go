package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/database"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/repository"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/rest"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
	"github.com/dmytrodemianchuk/go-auth-mongo/pkg/hash"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")

	fmt.Println("JWT Secret:", jwtSecret)

	// Establish a MongoDB connection
	db, err := database.NewMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	booksRepo := repository.NewBooksRepository(db)
	usersRepo := repository.NewUsersRepository(db)

	// Initialize password hasher
	passwordHasher := hash.NewHasher() // Assuming you have a simple hasher implementation

	// Initialize services
	booksService := service.NewBooks(booksRepo)
	usersService := service.NewUsers(usersRepo, passwordHasher, []byte(jwtSecret), 24*time.Hour)

	// Initialize handlers
	handler := rest.NewHandler(booksService, usersService)
	mux := handler.InitRouter()

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is running on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed:", err)
	}
}
