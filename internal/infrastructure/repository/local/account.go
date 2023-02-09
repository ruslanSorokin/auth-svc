package local

import (
	"context"
	"sync"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"
)

// AccountRepository is in-memory implementation of IAccountRepository
type AccountRepository struct {
	db map[string]*model.Account
	m  *sync.RWMutex
}

var _ repository.ISessionRepository = (*SessionRepository)(nil)

// NewAccountRepository is a default constructor
func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		db: make(map[string]*model.Account),
		m:  new(sync.RWMutex),
	}
}

// GetAccountByID returns Account model by ID
func (r *AccountRepository) GetAccountByID(ctx context.Context, id string) (*model.Account, error) {
	r.m.RLock()
	s := *r.db[id]
	r.m.RUnlock()

	return &s, nil
}

// GetAccountByEmail returns Account model by email
func (r *AccountRepository) GetAccountByEmail(ctx context.Context, eMail string) (*model.Account, error) {
	var s model.Account

	r.m.RLock()
	for _, v := range r.db {
		if v.EMail == eMail {
			s = *v
			break
		}
	}
	r.m.RUnlock()

	return &s, nil
}

// GetAccountByLogin returns Account model by login
func (r *AccountRepository) GetAccountByLogin(ctx context.Context, login string) (*model.Account, error) {
	var s model.Account

	r.m.RLock()
	for _, v := range r.db {
		if v.Login == login {
			s = *v
			break
		}
	}
	r.m.RUnlock()

	return &s, nil
}
