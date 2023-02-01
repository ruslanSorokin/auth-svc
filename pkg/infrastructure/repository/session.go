package repository

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
)

// ISessionRepository is an interface with custom CRUD-like operations for Session entity.
type ISessionRepository interface {
	CreateSession(ctx context.Context, s *model.Session) (string, error)
	DeleteSessionByID(ctx context.Context, ID string) error
	DeleteSessionByUserID(ctx context.Context, userID string) error
	UpdateSessionByID(ctx context.Context, ID string, s *model.Session) error
	UpdateSessionByUserID(ctx context.Context, userID string, s *model.Session) error
	GetSessionByID(ctx context.Context, ID string) (*model.Session, error)
	GetSessionByUserID(ctx context.Context, userID string) (*model.Session, error)
}
