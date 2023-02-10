package service

import (
	"github.com/ruslanSorokin/authentication-service/pkg/app/orchestrator/service"
	"github.com/ruslanSorokin/authentication-service/pkg/infra/repository"

	"github.com/ruslanSorokin/authentication-service/cmd/orchestrator/config"
)

type Orchestrator struct {
	cfg          *config.Service
	sessionStore *repository.ISessionStore
	accountStore *repository.IAccountStore
}

var _ service.IOrchestrator = (*Orchestrator)(nil)

func New(c *config.Service, s *repository.ISessionStore, a *repository.IAccountStore) *Orchestrator {
	return &Orchestrator{
		cfg:          c,
		sessionStore: s,
		accountStore: a,
	}
}
