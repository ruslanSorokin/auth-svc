package model

import (
	"time"
)

// Header of the both Access & Refresh tokens.
type Header struct {
	Algorithm string    `json:"alg"`
	Type      string    `json:"typ"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
}

// Payload of the Access Token.
type Payload struct {
	UserID string
}

// AccessToken used to store information for user's session.
type AccessToken struct {
	Header  Header
	Payload Payload
}

// RefreshToken used to get a new token pair.
type RefreshToken struct {
	Header Header
}

// TokenPair is a pair of both Access and Refresh tokens.
type TokenPair struct {
	AToken AccessToken
	RToken RefreshToken
}
