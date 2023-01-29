package mongo

import (
	"context"

	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"
	"github.com/ruslanSorokin/auth-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository is MongoDB implementation of IUserRepository
type UserRepository struct {
	db *mongo.Collection
}

var _ repository.IUserRepository = (*UserRepository)(nil)

// NewUserRepository is a default constructor
func NewUserRepository(URI, dbName, collectionName string) *UserRepository {
	return &UserRepository{
		db: newInstance(URI, dbName).Collection(collectionName),
	}
}

// GetUserByID returns User model by UserID
func (r *UserRepository) GetUserByID(ctx context.Context, userID *string) (*model.User, error) {
	s := new(model.User)

	res := r.db.FindOne(ctx, bson.M{
		"id": userID,
	})
	if res.Err() != nil {
		return nil, repository.ErrUserNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetUserByEmail returns User model by UserEmail
func (r *UserRepository) GetUserByEmail(ctx context.Context, userEmail *string) (*model.User, error) {
	s := new(model.User)

	res := r.db.FindOne(ctx, bson.M{
		"email": userEmail,
	})
	if res.Err() != nil {
		return nil, repository.ErrUserNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
