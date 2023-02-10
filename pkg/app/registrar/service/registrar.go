package service

import "context"

type IRegistrar interface {
	SignUp(ctx context.Context, email, login, password string) error
}
