package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/model"
)

// IUserRepoEngine is an interface with only custom CRUD-like operations for User entity.
type IUserRepoEngine interface {
	GetUserById(ctx context.Context, uid int) (user model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
}
