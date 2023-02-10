package repository

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
)

// IAccountStore is an interface with custom CRUD-like operations for User entity.
type IAccountStore interface {
	GetAccountByID(ctx context.Context, id string) (*model.Account, error)
	GetAccountByEmail(ctx context.Context, eMail string) (*model.Account, error)
	GetAccountByLogin(ctx context.Context, login string) (*model.Account, error)
}
