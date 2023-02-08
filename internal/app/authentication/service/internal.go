package service

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/app/authentication/service"

	"github.com/ruslanSorokin/auth-service/cmd/authentication/config"
)

type InternalService struct {
	cfg *config.Service
}

var _ service.IInternalService = (*InternalService)(nil)

func NewInternalService(c *config.Service) *InternalService {
	return &InternalService{
		cfg: c,
	}
}

func (s *InternalService) GetSecrets(ctx context.Context, serviceName string, serviceGUID string, password string) (*[3]string, error) {
	panic("not implemented") // TODO: Implement
}

func (s *InternalService) ForceLogoutAll(ctx context.Context, serviceName string, serviceGUID string, password string) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (s *InternalService) ForceLogout(ctx context.Context, serviceName string, serviceGUID string, password string) error {
	panic("not implemented") // TODO: Implement
}
