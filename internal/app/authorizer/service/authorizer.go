package service

import (
	"github.com/ruslanSorokin/authentication-service/pkg/app/authorizer/service"
	"github.com/ruslanSorokin/authentication-service/pkg/infra/repository"

	"github.com/ruslanSorokin/authentication-service/cmd/authorizer/config"
)

type Authorizer struct {
	cfg          *config.Service
	sessionStore *repository.ISessionStore
	accountStore *repository.IAccountStore
}

var _ service.IAuthorizer = (*Authorizer)(nil)

func New(c *config.Service, s *repository.ISessionStore, a *repository.IAccountStore) *Authorizer {
	return &Authorizer{
		cfg:          c,
		sessionStore: s,
		accountStore: a,
	}
}
