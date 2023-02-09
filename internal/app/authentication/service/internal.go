package service

import (
	"context"

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

func (a *InternalAuthorizer) GetSecrets(ctx context.Context, serviceName string, serviceGUID string, password string) (*[3]string, error) {
	panic("not implemented") // TODO: Implement
}

func (a *InternalAuthorizer) ForceLogoutAll(ctx context.Context, serviceName string, serviceGUID string, password string) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (a *InternalAuthorizer) ForceLogout(ctx context.Context, serviceName string, serviceGUID string, password string) error {
	panic("not implemented") // TODO: Implement
}
