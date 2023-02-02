package model

import "time"

// Session stored in the authentication service's DataBase.
type Session struct {
	ID          string    `bson:"id" json:"id"`
	RToken      string    `bson:"r_token" json:"r_token"`
	UserID      string    `bson:"user_id" json:"user_id"`
	Fingerprint string    `bson:"fingerprint" json:"fingerprint"`
	IP          string    `bson:"ip" json:"ip"`
	ExpiresAt   time.Time `bson:"expires_at" json:"expires_at"`
}
