package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/model"
)

// IUserRepository is an interface with custom CRUD-like operations for User entity.
type IUserRepository interface {
	GetUserByID(ctx context.Context, userID *string) (*model.User, error)
	GetUserByEmail(ctx context.Context, userEmail *string) (*model.User, error)
}
