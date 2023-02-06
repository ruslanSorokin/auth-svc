package mongo

import (
	"context"

	"github.com/ruslanSorokin/auth-service/internal/app/config"
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
		db: newInstance(cfg.Mongo.URI).
			Database(cfg.Mongo.DBName).
			Collection(cfg.Mongo.TableName.Session),
	}
}

// CreateSession creates new session with given Session model and returns its ID
func (r *SessionRepository) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	res, err := r.db.InsertOne(ctx, s)
	if err != nil {
		return "", err
	}

	id := string(res.InsertedID.(primitive.ObjectID).Hex())
	s.ID = id
	return id, nil
}

// DeleteSessionByID deletes session by ID
func (r *SessionRepository) DeleteSessionByID(ctx context.Context, ID string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{
		"id": ID,
	})

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// DeleteSessionByUserID deletes session by UserID
func (r *SessionRepository) DeleteSessionByUserID(ctx context.Context, userID string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{
		"user_id": userID,
	})

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// UpdateSessionByID updates session by ID
func (r *SessionRepository) UpdateSessionByID(ctx context.Context, ID string, s *model.Session) error {
	res, err := r.db.UpdateOne(ctx, bson.M{
		"id": ID,
	}, s)

	if err != nil {
		return err
	}

	if res.UpsertedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// UpdateSessionByUserID updates session by UserID
func (r *SessionRepository) UpdateSessionByUserID(ctx context.Context, userID string, s *model.Session) error {
	res, err := r.db.UpdateOne(ctx, bson.M{
		"user_id": userID,
	}, s)

	if err != nil {
		return err
	}

	if res.UpsertedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// GetSessionByID returns Session model by ID
func (r *SessionRepository) GetSessionByID(ctx context.Context, ID string) (*model.Session, error) {
	s := new(model.Session)

	res := r.db.FindOne(ctx, bson.M{
		"id": ID,
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

// GetSessionByUserID returns Sessions model by UserID
func (r *SessionRepository) GetSessionByUserID(ctx context.Context, userID string) (*model.Session, error) {
	s := new(model.Session)

	res := r.db.FindOne(ctx, bson.M{
		"user_id": userID,
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
