package model

type userID uint64

type User struct {
	UserId   userID
	Username string
	Email    string
	Password string
}
