package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/model"
)

// ISessionRepoEngine is an interface with only custom CRUD-like operations for User entity.
type ISessionRepoEngine interface {
	DropUserRefreshTokenByUserID(ctx context.Context, userID int) (err error)
	PatchUserRefreshTokenByUserID(ctx context.Context, userID int, rToken model.RefreshToken) (err error)
}
