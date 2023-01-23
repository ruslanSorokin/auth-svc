package token

import "time"

type Header struct {
	Algorithm string    `json:"alg"`
	Type      string    `json:"typ"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
}

type Payload struct {
	UserID uint64
}
