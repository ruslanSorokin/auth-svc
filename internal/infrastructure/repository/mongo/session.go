package mongo

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/cmd/authentication/config"
	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infrastructure/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SessionStore is MongoDB implementation of ISessionRepository
type SessionStore struct {
	db *mongo.Collection
}

var _ repository.ISessionStore = (*SessionStore)(nil)

// NewSessionStore is a default constructor
func NewSessionStore(URI, db, collection string) *SessionStore {
	return &SessionStore{
		db: newInstance(URI).
			Database(db).
			Collection(collection),
	}
}

// NewSessionStoreFromConfig is custom constructor from config
func NewSessionStoreFromConfig(cfg *config.DB) *SessionStore {
	return &SessionStore{
		db: newInstance(cfg.Mongo.Session.URI).
			Database(cfg.Mongo.Session.DB).
			Collection(cfg.Mongo.Session.Table),
	}
}

// CreateSession creates new session with given Session model and returns its ID
func (r *SessionStore) CreateSession(ctx context.Context, s *model.Session) (string, error) {
	res, err := r.db.InsertOne(ctx, s)
	if err != nil {
		return "", err
	}

	id := string(res.InsertedID.(primitive.ObjectID).Hex())
	s.ID = id
	return id, nil
}

// DeleteSessionByID deletes session by ID
func (r *SessionStore) DeleteSessionByID(ctx context.Context, id string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{
		"id": id,
	})

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// DeleteSessionByAccountID deletes session by AccountID
func (r *SessionStore) DeleteSessionByAccountID(ctx context.Context, accountID string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{
		"account_id": accountID,
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
func (r *SessionStore) UpdateSessionByID(ctx context.Context, id string, s *model.Session) error {
	res, err := r.db.UpdateOne(ctx, bson.M{
		"id": id,
	}, s)

	if err != nil {
		return err
	}

	if res.UpsertedCount != 1 {
		return repository.ErrSessionNotFound
	}
	return nil
}

// UpdateSessionByAccountID updates session by AccountID
func (r *SessionStore) UpdateSessionByAccountID(ctx context.Context, accountID string, s *model.Session) error {
	res, err := r.db.UpdateOne(ctx, bson.M{
		"account_id": accountID,
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
func (r *SessionStore) GetSessionByID(ctx context.Context, id string) (*model.Session, error) {
	s := new(model.Session)

	res := r.db.FindOne(ctx, bson.M{
		"id": id,
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

// GetSessionByAccountID returns Sessions model by AccountID
func (r *SessionStore) GetSessionByAccountID(ctx context.Context, accountID string) (*model.Session, error) {
	s := new(model.Session)

	res := r.db.FindOne(ctx, bson.M{
		"account_id": accountID,
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
