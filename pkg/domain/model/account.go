package model

import "time"

// Account stores information needed for authentication
type Account struct {
	ID             string    `bson:"id"`
	Login          string    `bson:"login"`
	EMail          string    `bson:"email"`
	SaltedPassword string    `bson:"salted_password"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
	DeletedAt      time.Time `bson:"deleted_at"`
}
