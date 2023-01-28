package mongo

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"
	"github.com/ruslanSorokin/auth-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository struct {
	db *mongo.Collection
}

var _ repository.ISessionRepository = SessionRepository{}

func NewSessionRepository(URI, dbName, collectionName string) *SessionRepository {
	return &SessionRepository{
		db: newInstance(URI, dbName).Collection(collectionName),
	}
}

func (r SessionRepository) CreateSession(ctx context.Context, s *model.Session) (*string, error) {
	res, err := r.db.InsertOne(ctx, s)
	id := string(res.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return &id, err
	}
	return &id, nil
}

func (r SessionRepository) DeleteSessionByID(ctx context.Context, ID *string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"id": *ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r SessionRepository) DeleteSessionByUserID(ctx context.Context, userID *string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{
		"user_id": userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r SessionRepository) UpdateSessionByID(ctx context.Context, ID *string, s *model.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"id": ID,
	}, s)
	if err != nil {
		return err
	}
	return nil
}
func (r SessionRepository) UpdateSessionByUserID(ctx context.Context, userID *string, s *model.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{
		"user_id": userID,
	}, s)
	if err != nil {
		return err
	}
	return nil
}

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
