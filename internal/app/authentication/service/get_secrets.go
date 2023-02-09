package service

import (
	"context"
)

func (a *InternalAuthorizer) GetSecrets(ctx context.Context, serviceName string, serviceGUID string, password string) (*[3]string, error) {
	panic("not implemented") // TODO: Implement
}
