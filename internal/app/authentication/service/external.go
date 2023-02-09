package service

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/app/authentication/service"
	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"

	"github.com/ruslanSorokin/authentication-service/cmd/authentication/config"
)

type ExternalAuthorizer struct {
	cfg         *config.Service
	sessionRepo repository.ISessionRepository
	userRepo    repository.IUserRepository
}

var _ service.IExternalAuthorizer = (*ExternalAuthorizer)(nil)

func NewExternalAuthorizer(c *config.Service, s repository.ISessionRepository, u repository.IUserRepository) *ExternalAuthorizer {
	return &ExternalAuthorizer{
		cfg:         c,
		sessionRepo: s,
		userRepo:    u,
	}
}

func (a *ExternalAuthorizer) Login(ctx context.Context, login, password string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}

func (a *ExternalAuthorizer) RefreshTokenPair(ctx context.Context, rToken string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}

func (a *ExternalAuthorizer) Logout(ctx context.Context, rToken string) error {
	panic("not implemented") // TODO: Implement
}

func (a *ExternalAuthorizer) LogoutAll(ctx context.Context, rToken string) error {
	panic("not implemented") // TODO: Implement
}
