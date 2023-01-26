package model

// User is an entity of user used to pass over business logic.
type User struct {
	UserID         string `json:"id"`
	Username       string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}
