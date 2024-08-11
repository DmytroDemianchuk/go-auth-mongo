package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
)

type BooksRepository struct {
	db *mongo.Database
}

func NewBooksRepository(db *mongo.Database) *BooksRepository {
	return &BooksRepository{db: db}
}

func (r *BooksRepository) Create(ctx context.Context, book domain.Book) error {
	_, err := r.db.Collection("books").InsertOne(ctx, book)
	return err
}

func (r *BooksRepository) GetByID(ctx context.Context, id string) (domain.Book, error) {
	var book domain.Book
	err := r.db.Collection("books").FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	return book, err
}

func (r *BooksRepository) GetAll(ctx context.Context) ([]domain.Book, error) {
	cursor, err := r.db.Collection("books").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []domain.Book
	for cursor.Next(ctx) {
		var book domain.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BooksRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Collection("books").DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *BooksRepository) Update(ctx context.Context, id string, inp domain.Book) error {
	_, err := r.db.Collection("books").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": inp})
	return err
}
