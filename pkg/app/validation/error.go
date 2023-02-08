package validation

import (
	"errors"
)

var (
	// ErrIncorrectUsernameLength returned if login doesn't fit into its validation parameters
	ErrIncorrectUsernameLength = errors.New("incorrect login length")
	// ErrIncorrectPasswordLength returned if password doesn't fit into its validation parameters
	ErrIncorrectPasswordLength = errors.New("incorrect password length")
)
