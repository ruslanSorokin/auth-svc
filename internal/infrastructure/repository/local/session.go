package local

import (
	"context"
	"math/rand"
	"sync"

	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"
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

const (
	keyLength = 32
)

func randomString(l int) string {
	r := make([]byte, l)
	for i := 0; i < l; i++ {
		r[i] = byte(rand.Intn(255))
	}
	return string(r)
}

// CreateSession creates new session with given Session model and returns its ID
func (r *SessionRepository) CreateSession(ctx context.Context, s *model.Session) (*string, error) {
	id := randomString(keyLength)

	r.m.Lock()
	defer r.m.Unlock()

	s.ID = id
	r.db[id] = s

	return &id, nil
}

// DeleteSessionByID deletes session by ID
func (r *SessionRepository) DeleteSessionByID(ctx context.Context, ID *string) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.db[*ID]
	if !ok {
		return repository.ErrSessionNotFound
	}

	delete(r.db, *ID)

	return nil
}

// DeleteSessionByUserID deletes session by UserID
func (r *SessionRepository) DeleteSessionByUserID(ctx context.Context, userID *string) error {
	r.m.Lock()
	defer r.m.Unlock()

	for k, v := range r.db {
		if v.UserID == *userID {
			delete(r.db, k)
			return nil
		}
	}

	return repository.ErrSessionNotFound
}

// UpdateSessionByID updates session by ID
func (r *SessionRepository) UpdateSessionByID(ctx context.Context, ID *string, s *model.Session) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.db[*ID]
	if !ok {
		return repository.ErrSessionNotFound
	}

	r.db[*ID] = s

	return nil
}

// UpdateSessionByUserID updates session by UserID
func (r *SessionRepository) UpdateSessionByUserID(ctx context.Context, userID *string, s *model.Session) error {
	r.m.Lock()
	for k, v := range r.db {
		if v.UserID == *userID {
			r.db[k] = s
			break
		}
	}
	r.m.Unlock()

	return nil
}

// GetSessionByID returns Session model by ID
func (r *SessionRepository) GetSessionByID(ctx context.Context, ID *string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	_, ok := r.db[*ID]
	if !ok {
		return nil, repository.ErrSessionNotFound
	}

	s := r.db[*ID]

	return s, nil
}

// GetSessionByUserID returns Sessions model by UserID
func (r *SessionRepository) GetSessionByUserID(ctx context.Context, userID *string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, v := range r.db {
		if v.UserID == *userID {
			return v, nil
		}
	}

	return nil, repository.ErrSessionNotFound
}
