package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
)

// IUserRepository is an interface with custom CRUD-like operations for User entity.
type IUserRepository interface {
	GetUserByGUID(ctx context.Context, GUID string) (*model.User, error)
	GetUserByEmail(ctx context.Context, Email string) (*model.User, error)
}
