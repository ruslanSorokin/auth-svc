package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-svc/pkg/domain/model"
)

// ISessionStore is an interface with custom CRUD-like operations for Session entity.
type ISessionStore interface {
	CreateSession(ctx context.Context, s *model.Session) (string, error)
	DeleteSessionByID(ctx context.Context, id string) error
	DeleteSessionByAccountID(ctx context.Context, accountID string) error
	UpdateSessionByID(ctx context.Context, id string, s *model.Session) error
	UpdateSessionByAccountID(ctx context.Context, accountID string, s *model.Session) error
	GetSessionByID(ctx context.Context, id string) (*model.Session, error)
	GetSessionByAccountID(ctx context.Context, accountID string) (*model.Session, error)
}
