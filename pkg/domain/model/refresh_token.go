package model

import "github.com/golang-jwt/jwt"

// RefreshToken doesn't store any authentication information about user's account
// and only being used to get new TokenPair.
type RefreshToken struct {
	jwt.StandardClaims
}
