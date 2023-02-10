package service

import (
	"context"
)

func (o *Orchestrator) GetSecrets(ctx context.Context, serviceName string, serviceGUID string, password string) (*[3]string, error) {
	panic("not implemented") // TODO: Implement
}
