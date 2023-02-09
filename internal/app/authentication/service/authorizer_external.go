package service

import (
	"github.com/ruslanSorokin/authentication-service/pkg/app/authentication/service"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"

	"github.com/ruslanSorokin/authentication-service/cmd/authentication/config"
)

type ExternalAuthorizer struct {
	cfg         *config.Service
	sessionRepo repository.ISessionRepository
	accountRepo repository.IAccountRepository
}

var _ service.IExternalAuthorizer = (*ExternalAuthorizer)(nil)

func NewExternalAuthorizer(c *config.Service, s repository.ISessionRepository, a repository.IAccountRepository) *ExternalAuthorizer {
	return &ExternalAuthorizer{
		cfg:         c,
		sessionRepo: s,
		accountRepo: a,
	}
}
