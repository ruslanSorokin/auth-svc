package service

import (
	"github.com/ruslanSorokin/auth-svc/pkg/app/registrar/service"
	"github.com/ruslanSorokin/auth-svc/pkg/infra/repository"

	"github.com/ruslanSorokin/auth-svc/cmd/registrar/config"
)

type Registrar struct {
	cfg          config.RegistrarConfig
	accountStore repository.IAccountStore
}

var _ service.IRegistrar = (*Registrar)(nil)

func New(c config.RegistrarConfig, a repository.IAccountStore) *Registrar {
	return &Registrar{
		cfg:          c,
		accountStore: a,
	}
}
