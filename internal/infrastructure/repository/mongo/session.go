package mongo

import (
	"context"

	"github.com/ruslanSorokin/auth-service/cmd/authentication/config"
	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SessionRepository is MongoDB implementation of ISessionRepository
type SessionRepository struct {
	db *mongo.Collection
}

var _ repository.ISessionRepository = (*SessionRepository)(nil)

// NewSessionRepository is a default constructor
func NewSessionRepository(URI, dbName, collectionName string) *SessionRepository {
	return &SessionRepository{
		db: newInstance(URI).
			Database(dbName).
			Collection(collectionName),
	}
}

// NewSessionRepositoryFromConfig is custom constructor from config
func NewSessionRepositoryFromConfig(cfg *config.DB) *SessionRepository {
	return &SessionRepository{
		db: newInstance(cfg.Mongo.Session.URI).
			Database(cfg.Mongo.Session.DBName).
			Collection(cfg.Mongo.Session.TableName),
	}
}

// CreateSession creates new session with given Session model and returns its GUID
func (r *SessionRepository) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	res, err := r.db.InsertOne(ctx, s)
	if err != nil {
		return "", err
	}

	id := string(res.InsertedID.(primitive.ObjectID).Hex())
	s.ID = id
	return id, nil
}

// DeleteSessionByGUID deletes session by GUID
func (r *SessionRepository) DeleteSessionByGUID(ctx context.Context, GUID string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{
		"id": GUID,
	})

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// DeleteSessionByUserGUID deletes session by UserGUID
func (r *SessionRepository) DeleteSessionByUserGUID(ctx context.Context, userGUID string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{
		"user_id": userGUID,
	})

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// UpdateSessionByGUID updates session by GUID
func (r *SessionRepository) UpdateSessionByGUID(ctx context.Context, GUID string, s *model.Session) error {
	res, err := r.db.UpdateOne(ctx, bson.M{
		"id": GUID,
	}, s)

	if err != nil {
		return err
	}

	if res.UpsertedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// UpdateSessionByUserGUID updates session by UserGUID
func (r *SessionRepository) UpdateSessionByUserGUID(ctx context.Context, userGUID string, s *model.Session) error {
	res, err := r.db.UpdateOne(ctx, bson.M{
		"user_id": userGUID,
	}, s)

	if err != nil {
		return err
	}

	if res.UpsertedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// GetSessionByGUID returns Session model by GUID
func (r *SessionRepository) GetSessionByGUID(ctx context.Context, GUID string) (*model.Session, error) {
	s := new(model.Session)

	res := r.db.FindOne(ctx, bson.M{
		"id": GUID,
	})
	if res.Err() != nil {
		return nil, repository.ErrSessionNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// GetSessionByUserGUID returns Sessions model by UserGUID
func (r *SessionRepository) GetSessionByUserGUID(ctx context.Context, userGUID string) (*model.Session, error) {
	s := new(model.Session)

	res := r.db.FindOne(ctx, bson.M{
		"user_id": userGUID,
	})
	if res.Err() != nil {
		return nil, repository.ErrSessionNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
