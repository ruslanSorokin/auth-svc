package model

import (
	"github.com/golang-jwt/jwt"
)

// AccessToken stores authentication information for user's account.
// Do not change the struct, instead add new fields to Payload struct.
type AccessToken struct {
	jwt.StandardClaims
	Payload
}
