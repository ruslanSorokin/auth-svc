package mongo

import (
	"context"

	"github.com/ruslanSorokin/authentication-service/cmd/registrar/config"
	"github.com/ruslanSorokin/authentication-service/pkg/domain/model"
	"github.com/ruslanSorokin/authentication-service/pkg/infra/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AccountRepository is MongoDB implementation of IAccountRepository
type AccountRepository struct {
	db *mongo.Collection
}

var _ repository.IAccountStore = (*AccountRepository)(nil)

// NewAccountStore is a default constructor
func NewAccountStore(URI, db, collection string) *AccountRepository {
	return &AccountRepository{
		db: newInstance(URI).
			Database(db).
			Collection(collection),
	}
}

// NewAccountStoreFromConfig is custom constructor from config
func NewAccountStoreFromConfig(cfg *config.Repository) *AccountRepository {
	return &AccountRepository{
		db: newInstance(cfg.Mongo.Account.URI).
			Database(cfg.Mongo.Account.DB).
			Collection(cfg.Mongo.Account.Table),
	}
}

// GetAccountByID returns Account model by id
func (r *AccountRepository) GetAccountByID(ctx context.Context, id string) (*model.Account, error) {
	s := new(model.Account)

	res := r.db.FindOne(ctx, bson.M{
		"id": id,
	})
	if res.Err() != nil {
		return nil, repository.ErrAccountNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetAccountByEmail returns Account model by email
func (r *AccountRepository) GetAccountByEmail(ctx context.Context, eMail string) (*model.Account, error) {
	s := new(model.Account)

	res := r.db.FindOne(ctx, bson.M{
		"email": eMail,
	})
	if res.Err() != nil {
		return nil, repository.ErrAccountNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetAccountByLogin returns Account model by login
func (r *AccountRepository) GetAccountByLogin(ctx context.Context, login string) (*model.Account, error) {
	s := new(model.Account)

	res := r.db.FindOne(ctx, bson.M{
		"login": login,
	})
	if res.Err() != nil {
		return nil, repository.ErrAccountNotFound
	}

	err := res.Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
