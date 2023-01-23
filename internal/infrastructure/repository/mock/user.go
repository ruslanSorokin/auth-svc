package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/internal/infrastructure/repository"
	. "github.com/ruslanSorokin/auth-service/internal/model"
)

type UserRepoEngine struct {
}

var _ repository.IUserRepoEngine = UserRepoEngine{}

func (r UserRepoEngine) GetUserById(ctx context.Context, uid int) (user User, err error) {
	return
}

func (r UserRepoEngine) GetUserByEmail(ctx context.Context, email string) (user User, err error) {
	return
}
