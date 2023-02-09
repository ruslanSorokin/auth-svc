package service

import (
	"github.com/ruslanSorokin/authentication-service/pkg/app/authentication/service"

	"github.com/ruslanSorokin/authentication-service/cmd/authentication/config"
)

type InternalAuthorizer struct {
	cfg *config.Service
}

var _ service.IInternalAuthorizer = (*InternalAuthorizer)(nil)

func NewInternalAuthorizer(c *config.Service) *InternalAuthorizer {
	return &InternalAuthorizer{
		cfg: c,
	}
}
