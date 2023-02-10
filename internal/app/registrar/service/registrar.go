package service

import (
	"github.com/ruslanSorokin/authentication-service/pkg/app/registrar/service"
	"github.com/ruslanSorokin/authentication-service/pkg/infra/repository"

	"github.com/ruslanSorokin/authentication-service/cmd/registrar/config"
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
