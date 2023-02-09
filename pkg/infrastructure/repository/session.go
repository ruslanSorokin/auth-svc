package repository

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
)

// ISessionRepository is an interface with custom CRUD-like operations for Session entity.
type ISessionRepository interface {
	CreateSession(ctx context.Context, s *model.Session) (string, error)
	DeleteSessionByGUID(ctx context.Context, GUID string) error
	DeleteSessionByUserGUID(ctx context.Context, userGUID string) error
	UpdateSessionByGUID(ctx context.Context, GUID string, s *model.Session) error
	UpdateSessionByUserGUID(ctx context.Context, userGUID string, s *model.Session) error
	GetSessionByGUID(ctx context.Context, GUID string) (*model.Session, error)
	GetSessionByUserGUID(ctx context.Context, userGUID string) (*model.Session, error)
}
