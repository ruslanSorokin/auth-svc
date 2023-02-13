package service

import (
	"context"

	"github.com/ruslanSorokin/auth-svc/pkg/domain/model"
)

type IAuthorizer interface {
	Login(ctx context.Context, login, password string) (*model.TokenPair, error)
	RefreshTokenPair(ctx context.Context, rToken string) (*model.TokenPair, error)
	Logout(ctx context.Context, rToken string) error
	LogoutAll(ctx context.Context, rToken string) error
}
