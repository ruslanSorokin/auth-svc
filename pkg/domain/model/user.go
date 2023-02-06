package model

// User is an entity of user used to pass over business logic.
type User struct {
	ID           string `bson:"id" json:"id,omitempty"`
	Login        string `bson:"login" json:"login,omitempty"`
	Email        string `bson:"email" json:"email,omitempty"`
	PasswordHash string `bson:"password_hash" json:"password_hash,omitempty"`
}
