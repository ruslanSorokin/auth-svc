package repository

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
)

// IUserRepository is an interface with custom CRUD-like operations for User entity.
type IUserRepository interface {
	GetUserByID(ctx context.Context, ID string) (*model.User, error)
	GetUserByEmail(ctx context.Context, Email string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
}
