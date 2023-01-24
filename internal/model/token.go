package model

import "github.com/ruslanSorokin/auth-service/internal/model/token"

type Payload struct {
	UserId userID
}

type AccessToken struct {
	token.Header
	Payload
}

type RefreshToken struct {
	token.Header
}

type TokenPair struct {
	AToken AccessToken
	RToken RefreshToken
}
