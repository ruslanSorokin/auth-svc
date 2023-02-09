package repository

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
)

// ISessionRepository is an interface with custom CRUD-like operations for Session entity.
type ISessionRepository interface {
	CreateSession(ctx context.Context, s *model.Session) (string, error)
	DeleteSessionByID(ctx context.Context, ID string) error
	DeleteSessionByAccountID(ctx context.Context, accountID string) error
	UpdateSessionByID(ctx context.Context, ID string, s *model.Session) error
	UpdateSessionByAccountID(ctx context.Context, accountID string, s *model.Session) error
	GetSessionByID(ctx context.Context, ID string) (*model.Session, error)
	GetSessionByAccountID(ctx context.Context, accountID string) (*model.Session, error)
}
