package service

import (
	"context"

	"github.com/ruslanSorokin/auth-svc/pkg/domain/model"
)

func (a *Authorizer) Login(ctx context.Context, login, password string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}
