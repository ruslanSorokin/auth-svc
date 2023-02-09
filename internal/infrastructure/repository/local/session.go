package local

import (
	"context"
	"sync"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"
)

// SessionRepository is in-memory implementation of ISessionRepository
type SessionRepository struct {
	db map[string]*model.Session
	m  *sync.RWMutex
}

var _ repository.ISessionRepository = (*SessionRepository)(nil)

// NewSessionRepository is a default constructor
func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		db: make(map[string]*model.Session),
		m:  new(sync.RWMutex),
	}
}

// CreateSession creates new session with given Session model and returns its ID
func (r *SessionRepository) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	id := Hash(s)

	r.m.Lock()
	defer r.m.Unlock()

	s.ID = id
	r.db[id] = s

	return id, nil
}

// DeleteSessionByID deletes session by ID
func (r *SessionRepository) DeleteSessionByID(ctx context.Context, id string) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.db[id]
	if !ok {
		return repository.ErrSessionNotFound
	}

	delete(r.db, id)

	return nil
}

// DeleteSessionByAccountID deletes session by AccountID
func (r *SessionRepository) DeleteSessionByAccountID(ctx context.Context, accountID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	for k, v := range r.db {
		if v.AccountID == accountID {
			delete(r.db, k)
			return nil
		}
	}

	return repository.ErrSessionNotFound
}

// UpdateSessionByID updates session by ID
func (r *SessionRepository) UpdateSessionByID(ctx context.Context, id string, s *model.Session) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.db[id]
	if !ok {
		return repository.ErrSessionNotFound
	}

	r.db[id] = s

	return nil
}

// UpdateSessionByAccountID updates session by AccountID
func (r *SessionRepository) UpdateSessionByAccountID(ctx context.Context, accountID string, s *model.Session) error {
	r.m.Lock()
	for k, v := range r.db {
		if v.AccountID == accountID {
			r.db[k] = s
			break
		}
	}
	r.m.Unlock()

	return nil
}

// GetSessionByID returns Session model by ID
func (r *SessionRepository) GetSessionByID(ctx context.Context, id string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	_, ok := r.db[id]
	if !ok {
		return nil, repository.ErrSessionNotFound
	}

	s := r.db[id]

	return s, nil
}

// GetSessionByAccountID returns Sessions model by AccountID
func (r *SessionRepository) GetSessionByAccountID(ctx context.Context, accountID string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, v := range r.db {
		if v.AccountID == accountID {
			return v, nil
		}
	}

	return nil, repository.ErrSessionNotFound
}
