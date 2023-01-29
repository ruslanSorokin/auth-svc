package repository

import "errors"

var (
	ErrSessionNotFound = errors.New("session is not found")
)

var (
	ErrUserNotFound = errors.New("user is not found")
)
