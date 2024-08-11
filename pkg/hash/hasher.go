package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// Hasher is a simple implementation of the PasswordHasher interface.
type Hasher struct{}

// NewHasher creates a new Hasher instance.
func NewHasher() *Hasher {
	return &Hasher{}
}

// Hash hashes the given password.
func (h *Hasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
