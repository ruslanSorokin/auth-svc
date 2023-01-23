package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/internal/infrastructure/repository"
	. "github.com/ruslanSorokin/auth-service/internal/model"
)

type SessionRepoEngine struct {
}

var _ repository.ISessionRepoEngine = SessionRepoEngine{}

func (r SessionRepoEngine) DropUserRefreshTokenByUserID(ctx context.Context, userID int) (err error) {
	return
}

func (r SessionRepoEngine) PatchUserRefreshTokenByUserID(ctx context.Context, userID int, rToken RefreshToken) (err error) {
	return
}
