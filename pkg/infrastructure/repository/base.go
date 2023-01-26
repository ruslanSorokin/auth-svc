package repository

import "context"

// Model is an entity you insert to DB or receive from DB.
type Model interface{}

// Filter used to find models from DB you need.
type Filter interface{}

// IBaseRepoEngine is an interface with basic CRUD operations, which must be inside all others specific repository engines.
type IBaseRepoEngine interface {
	CreateOne(ctx context.Context, model Model) (id string, err error)
	ReadOne(ctx context.Context, filter Filter) (m Model, err error)
	UpdateOne(ctx context.Context, filter Filter, model Model) (err error)
	DeleteOne(ctx context.Context, filter Filter) (id int64, err error)
}
