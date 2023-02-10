package local

import (
	"context"
	"sync"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"
)

// AccountStore is in-memory implementation of IAccountRepository
type AccountStore struct {
	db map[string]*model.Account
	m  *sync.RWMutex
}

var _ repository.IAccountStore = (*AccountStore)(nil)

// NewAccountStore is a default constructor
func NewAccountStore() *AccountStore {
	return &AccountStore{
		db: make(map[string]*model.Account),
		m:  new(sync.RWMutex),
	}
}

// GetAccountByID returns Account model by ID
func (r *AccountStore) GetAccountByID(ctx context.Context, id string) (*model.Account, error) {
	r.m.RLock()
	s := *r.db[id]
	r.m.RUnlock()

	return &s, nil
}

// GetAccountByEmail returns Account model by email
func (r *AccountStore) GetAccountByEmail(ctx context.Context, eMail string) (*model.Account, error) {
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
func (r *AccountStore) GetAccountByLogin(ctx context.Context, login string) (*model.Account, error) {
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
