package domain

import "errors"

// ErrUserNotFound is returned when a user cannot be found.
var ErrUserNotFound = errors.New("user not found")
