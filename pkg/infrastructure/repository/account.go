package repository

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
)

// IAccountRepository is an interface with custom CRUD-like operations for User entity.
type IAccountRepository interface {
	GetAccountByID(ctx context.Context, ID string) (*model.Account, error)
	GetAccountByEmail(ctx context.Context, Email string) (*model.Account, error)
	GetAccountByLogin(ctx context.Context, login string) (*model.Account, error)
}
