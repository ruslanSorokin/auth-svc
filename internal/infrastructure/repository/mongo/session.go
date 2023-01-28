package mongo

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"
	"github.com/ruslanSorokin/auth-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SessionRepository is MongoDB implementation of ISessionRepository
type SessionRepository struct {
	db *mongo.Collection
}

var _ repository.ISessionRepository = SessionRepository{}

// NewSessionRepository is a default constructor
func NewSessionRepository(URI, dbName, collectionName string) *SessionRepository {
	return &SessionRepository{
		db: newInstance(URI, dbName).Collection(collectionName),
	}
}

// CreateSession creates new session with given Session model and returns its ID
func (r SessionRepository) CreateSession(ctx context.Context, s *model.Session) (*string, error) {
	res, err := r.db.InsertOne(ctx, s)
	id := string(res.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return &id, err
	}
	return &id, nil
}

// DeleteSessionByID deletes session by ID
func (r SessionRepository) DeleteSessionByID(ctx context.Context, ID *string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"id": *ID,
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteSessionByUserID deletes session by UserID
func (r SessionRepository) DeleteSessionByUserID(ctx context.Context, userID *string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"user_id": userID,
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateSessionByID updates session by ID
func (r SessionRepository) UpdateSessionByID(ctx context.Context, ID *string, s *model.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"id": ID,
	}, s)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSessionByUserID updates session by UserID
func (r SessionRepository) UpdateSessionByUserID(ctx context.Context, userID *string, s *model.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"user_id": userID,
	}, s)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionByID returns Session model by ID
func (r SessionRepository) GetSessionByID(ctx context.Context, ID *string) (*model.Session, error) {
	s := new(model.Session)
	res := r.db.FindOne(ctx, bson.M{
		"id": ID,
	})
	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetSessionByUserID returns Sessions model by UserID
func (r SessionRepository) GetSessionByUserID(ctx context.Context, userID *string) (*model.Session, error) {
	s := new(model.Session)
	res := r.db.FindOne(ctx, bson.M{
		"user_id": userID,
	})
	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
