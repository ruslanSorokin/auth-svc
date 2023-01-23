package repository

import (
	"context"

	. "github.com/ruslanSorokin/auth-service/internal/model"
)

type IUserRepoEngine interface {
	GetUserById(ctx context.Context, uid int) (user User, err error)
	GetUserByEmail(ctx context.Context, email string) (user User, err error)
}
