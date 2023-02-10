package service

import (
	"context"
)

type IOrchestrator interface {
	GetSecrets(ctx context.Context, serviceName string, serviceGUID string, password string) (*[3]string, error)
	ForceLogoutAll(ctx context.Context, serviceName string, serviceGUID string, password string) (int, error)
	ForceLogout(ctx context.Context, serviceName string, serviceGUID string, password string) error
}
