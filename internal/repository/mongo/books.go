package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BooksRepository handles operations on the books collection.
type BooksRepository struct {
	db *mongo.Database
}

// NewBooks creates a new BooksRepository.
func NewBooks(db *mongo.Database) *BooksRepository {
	return &BooksRepository{db: db}
}

// CreateBook adds a new book to the collection.
func (r *BooksRepository) CreateBook(ctx context.Context, book interface{}) error {
	collection := r.db.Collection("books")
	_, err := collection.InsertOne(ctx, book)
	return err
}

// GetBooks retrieves all books from the collection.
func (r *BooksRepository) GetBooks(ctx context.Context) ([]interface{}, error) {
	collection := r.db.Collection("books")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []interface{}
	for cursor.Next(ctx) {
		var book bson.M
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
