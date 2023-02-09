package model

import "time"

// Session stores information about active user's sessions
type Session struct {
	ID          string    `bson:"id"`
	RToken      string    `bson:"r_token"`
	AccountID   string    `bson:"account_id"`
	Fingerprint string    `bson:"fingerprint"`
	IP          string    `bson:"ip"`
	ExpiresAt   time.Time `bson:"expires_at"`
}
