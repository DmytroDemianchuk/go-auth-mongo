package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/repository/mongo"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/transport/rest"
	"github.com/dmytrodemianchuk/go-auth-mongo/pkg/database"
)

func main() {
	// Initialize MongoDB connection
	db, err := database.NewMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize dependencies
	booksRepo := mongo.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	// Initialize and run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
