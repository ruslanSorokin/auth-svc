package model

import "github.com/ruslanSorokin/auth-service/internal/model/token"

type AccessToken struct {
	token.Header
	token.Payload
}
