package model

const SaltLength = 16

// User is an entity of user used to pass over business logic.
type User struct {
	ID             string `bson:"id" json:"id,omitempty"`
	Login          string `bson:"login" json:"login,omitempty"`
	Email          string `bson:"email" json:"email,omitempty"`
	SaltedPassword string `bson:"salted_password" json:"salted_password,omitempty"`
}
