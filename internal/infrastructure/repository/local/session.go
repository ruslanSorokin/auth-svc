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

// CreateSession creates new session with given Session model and returns its GUID
func (r *SessionRepository) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	id := randomString(keyLength)

	r.m.Lock()
	defer r.m.Unlock()

	s.ID = id
	r.db[id] = s

	return id, nil
}

// DeleteSessionByGUID deletes session by GUID
func (r *SessionRepository) DeleteSessionByGUID(ctx context.Context, GUID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.db[GUID]
	if !ok {
		return repository.ErrSessionNotFound
	}

	delete(r.db, GUID)

	return nil
}

// DeleteSessionByUserGUID deletes session by UserGUID
func (r *SessionRepository) DeleteSessionByUserGUID(ctx context.Context, userGUID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	for k, v := range r.db {
		if v.UserID == userGUID {
			delete(r.db, k)
			return nil
		}
	}

	return repository.ErrSessionNotFound
}

// UpdateSessionByGUID updates session by GUID
func (r *SessionRepository) UpdateSessionByGUID(ctx context.Context, GUID string, s *model.Session) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.db[GUID]
	if !ok {
		return repository.ErrSessionNotFound
	}

	r.db[GUID] = s

	return nil
}

// UpdateSessionByUserGUID updates session by UserGUID
func (r *SessionRepository) UpdateSessionByUserGUID(ctx context.Context, userGUID string, s *model.Session) error {
	r.m.Lock()
	for k, v := range r.db {
		if v.UserID == userGUID {
			r.db[k] = s
			break
		}
	}
	r.m.Unlock()

	return nil
}

// GetSessionByGUID returns Session model by GUID
func (r *SessionRepository) GetSessionByGUID(ctx context.Context, GUID string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	_, ok := r.db[GUID]
	if !ok {
		return nil, repository.ErrSessionNotFound
	}

	s := r.db[GUID]

	return s, nil
}

// GetSessionByUserGUID returns Sessions model by UserGUID
func (r *SessionRepository) GetSessionByUserGUID(ctx context.Context, userGUID string) (*model.Session, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, v := range r.db {
		if v.UserID == userGUID {
			return v, nil
		}
	}

	return nil, repository.ErrSessionNotFound
}
