package service

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/app/authentication/service"
	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"

	"github.com/ruslanSorokin/auth-service/cmd/authentication/config"
)

type ExternalService struct {
	cfg         *config.Service
	sessionRepo repository.ISessionRepository
	userRepo    repository.IUserRepository
}

var _ service.IExternalService = (*ExternalService)(nil)

func NewExternalService(c *config.Service, s repository.ISessionRepository, u repository.IUserRepository) *ExternalService {
	return &ExternalService{
		cfg:         c,
		sessionRepo: s,
		userRepo:    u,
	}
}

func (s *ExternalService) Login(ctx context.Context, login, password string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ExternalService) RefreshTokenPair(ctx context.Context, rToken string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ExternalService) Logout(ctx context.Context, rToken string) error {
	panic("not implemented") // TODO: Implement
}

func (s *ExternalService) LogoutAll(ctx context.Context, rToken string) error {
	panic("not implemented") // TODO: Implement
}
