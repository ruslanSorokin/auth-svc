package mongo

import (
	"context"

	"github.com/ruslanSorokin/auth-service/cmd/registration/config"
	"github.com/ruslanSorokin/auth-service/pkg/domain/model"
	"github.com/ruslanSorokin/auth-service/pkg/infrastructure/repository"

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
		db: newInstance(URI).
			Database(dbName).
			Collection(collectionName),
	}
}

// NewUserRepositoryFromConfig is custom constructor from config
func NewUserRepositoryFromConfig(cfg *config.DB) *UserRepository {
	return &UserRepository{
		db: newInstance(cfg.Mongo.User.URI).
			Database(cfg.Mongo.User.DBName).
			Collection(cfg.Mongo.User.TableName),
	}
}

// GetUserByGUID returns User model by GUID
func (r *UserRepository) GetUserByGUID(ctx context.Context, GUID string) (*model.User, error) {
	s := new(model.User)

	res := r.db.FindOne(ctx, bson.M{
		"id": GUID,
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

// GetUserByEmail returns User model by email
func (r *UserRepository) GetUserByEmail(ctx context.Context, Email string) (*model.User, error) {
	s := new(model.User)

	res := r.db.FindOne(ctx, bson.M{
		"email": Email,
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

// GetUserByLogin returns User model by login
func (r *UserRepository) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	s := new(model.User)

	res := r.db.FindOne(ctx, bson.M{
		"login": login,
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
