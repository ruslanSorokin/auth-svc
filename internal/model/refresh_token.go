package model

import (
	"github.com/ruslanSorokin/auth-service/internal/model/token"
)

type RefreshToken struct {
	token.Header
}
