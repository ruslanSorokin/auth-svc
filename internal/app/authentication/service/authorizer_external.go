package service

import (
	"github.com/ruslanSorokin/authentication-service/pkg/app/authentication/service"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"

	"github.com/ruslanSorokin/authentication-service/cmd/authentication/config"
)

type ExternalAuthorizer struct {
	cfg          *config.Service
	sessionStore repository.ISessionStore
	accountStore repository.IAccountStore
}

var _ service.IExternalAuthorizer = (*ExternalAuthorizer)(nil)

func NewExternalAuthorizer(c *config.Service, s repository.ISessionStore, a repository.IAccountStore) *ExternalAuthorizer {
	return &ExternalAuthorizer{
		cfg:          c,
		sessionStore: s,
		accountStore: a,
	}
}
