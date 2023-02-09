package service

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
)

func (a *ExternalAuthorizer) RefreshTokenPair(ctx context.Context, rToken string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}
