package service

import (
	"context"

	"github.com/ruslanSorokin/auth-svc/pkg/domain/model"
)

func (a *Authorizer) RefreshTokenPair(ctx context.Context, rToken string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}
