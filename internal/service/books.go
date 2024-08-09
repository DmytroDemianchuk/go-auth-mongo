package service

import (
	"context"
	"fmt"
)

// BooksRepository defines the interface for book operations.
type BooksRepository interface {
	CreateBook(ctx context.Context, book interface{}) error
	GetBooks(ctx context.Context) ([]interface{}, error)
}

// BooksService provides methods to interact with book data.
type BooksService struct {
	repo BooksRepository
}

// NewBooks creates a new BooksService.
func NewBooks(repo BooksRepository) *BooksService {
	return &BooksService{repo: repo}
}

// AddBook adds a new book to the repository.
func (s *BooksService) AddBook(ctx context.Context, book interface{}) error {
	if err := s.repo.CreateBook(ctx, book); err != nil {
		return fmt.Errorf("failed to add book: %w", err)
	}
	return nil
}

// ListBooks retrieves all books from the repository.
func (s *BooksService) ListBooks(ctx context.Context) ([]interface{}, error) {
	books, err := s.repo.GetBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list books: %w", err)
	}
	return books, nil
}
