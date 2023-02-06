package local

import (
	"context"
	"sync"

	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"
)

// UserRepository is in-memory implementation of IUserRepository
type UserRepository struct {
	db map[string]*model.User
	m  *sync.RWMutex
}

var _ repository.ISessionRepository = (*SessionRepository)(nil)

// NewUserRepository is a default constructor
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: make(map[string]*model.User),
		m:  new(sync.RWMutex),
	}
}

// GetUserByGUID returns User model by GUID
func (r *UserRepository) GetUserByGUID(ctx context.Context, GUID string) (*model.User, error) {
	r.m.RLock()
	s := *r.db[GUID]
	r.m.RUnlock()

	return &s, nil
}

// GetUserByEmail returns User model by email
func (r *UserRepository) GetUserByEmail(ctx context.Context, Email string) (*model.User, error) {
	var s model.User

	r.m.RLock()
	for _, v := range r.db {
		if v.Email == Email {
			s = *v
			break
		}
	}
	r.m.RUnlock()

	return &s, nil
}

// GetUserByLogin returns User model by login
func (r *UserRepository) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	var s model.User

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
