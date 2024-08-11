package domain

import "time"

// Book represents the structure of a book in the system.
type Book struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Author      string    `json:"author" bson:"author"`
	PublishDate time.Time `json:"publish_date" bson:"publish_date"`
}

// SignUpInput represents the input for user signup.
type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInInput represents the input for user login.
type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User represents the structure of a user in the system.
type User struct {
	ID           string    `bson:"_id,omitempty"` // Use string for MongoDB ObjectID
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	Password     string    `bson:"password"`
	RegisteredAt time.Time `bson:"registered_at"`
}
