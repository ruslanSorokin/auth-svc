package model

// User is an entity of user used to pass over business logic.
type User struct {
	ID           string `bson:"id" json:"id,omitempty"`
	Username     string `bson:"username" json:"username,omitempty"`
	Email        string `bson:"email" json:"email,omitempty"`
	PasswordHash string `bson:"password_hash" json:"password_hash,omitempty"`
}
