package model

import (
	"time"
)

type Header struct {
	Algorithm string    `json:"alg"`
	Type      string    `json:"typ"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
}

type Payload struct {
	UserId userID
}

type AccessToken struct {
	Header  Header
	Payload Payload
}

type RefreshToken struct {
	Header
}

type TokenPair struct {
	AToken AccessToken
	RToken RefreshToken
}
