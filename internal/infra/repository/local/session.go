package local

import (
	"context"
	"sync"

	"github.com/ruslanSorokin/auth-svc/pkg/domain/model"
	"github.com/ruslanSorokin/auth-svc/pkg/infra/repository"
)

// SessionStore is in-memory implementation of ISessionRepository
type SessionStore struct {
	db map[string]*model.Session
	m  *sync.RWMutex
}

var _ repository.ISessionStore = (*SessionStore)(nil)

// NewSessionStore is a default constructor
func NewSessionStore() *SessionStore {
	return &SessionStore{
		db: make(map[string]*model.Session),
		m:  new(sync.RWMutex),
	}
}

// CreateSession creates new session with given Session model and returns its ID
func (r *SessionStore) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	id := Hash(s)

	r.m.Lock()
	defer r.m.Unlock()

	s.ID = id
	r.db[id] = s

	return id, nil
}

// DeleteSessionByID deletes session by ID
func (r *SessionStore) DeleteSessionByID(ctx context.Context, id string) error {
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
func (r *SessionStore) DeleteSessionByAccountID(ctx context.Context, accountID string) error {
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
func (r *SessionStore) UpdateSessionByID(ctx context.Context, id string, s *model.Session) error {
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
func (r *SessionStore) UpdateSessionByAccountID(ctx context.Context, accountID string, s *model.Session) error {
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
func (r *SessionStore) GetSessionByID(ctx context.Context, id string) (*model.Session, error) {
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
func (r *SessionStore) GetSessionByAccountID(ctx context.Context, accountID string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, v := range r.db {
		if v.AccountID == accountID {
			return v, nil
		}
	}

	return nil, repository.ErrSessionNotFound
}
