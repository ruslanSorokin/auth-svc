package model

// User is an entity of user used to pass over business logic.
type User struct {
	UserID         string `bson:"id" json:"id"`
	Username       string `bson:"name" json:"name"`
	Email          string `bson:"email" json:"email"`
	HashedPassword string `bson:"hashed_password" json:"hashed_password"`
}
