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

// GetUserByID returns User model by UserID
func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	r.m.RLock()
	s := *r.db[userID]
	r.m.RUnlock()

	return &s, nil
}

// GetUserByEmail returns User model by UserEmail
func (r *UserRepository) GetUserByEmail(ctx context.Context, userEmail string) (*model.User, error) {
	var s model.User

	r.m.RLock()
	for _, v := range r.db {
		if v.Email == userEmail {
			s = *v
			break
		}
	}
	r.m.RUnlock()

	return &s, nil
}
