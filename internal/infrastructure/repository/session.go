package repository

import (
	"context"

	. "github.com/ruslanSorokin/auth-service/internal/model"
)

type ISessionRepoEngine interface {
	DropUserRefreshTokenByUserID(ctx context.Context, userID int) (err error)
	PatchUserRefreshTokenByUserID(ctx context.Context, userID int, rToken RefreshToken) (err error)
}
