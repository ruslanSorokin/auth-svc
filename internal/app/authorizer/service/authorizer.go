package service

import (
	"github.com/ruslanSorokin/auth-svc/pkg/app/authorizer/service"
	"github.com/ruslanSorokin/auth-svc/pkg/infra/repository"

	"github.com/ruslanSorokin/auth-svc/cmd/authorizer/config"
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
