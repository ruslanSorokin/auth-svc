package model

// Session stored in the authentication service's DataBase.
type Session struct {
	ID          string `bson:"id" json:"id"`
	RToken      string `bson:"r_token" json:"r_token"`
	UserID      string `bson:"user_id" json:"user_id"`
	Fingerprint string `bson:"fingerprint" json:"fingerprint"`
	IP          string `bson:"ip" json:"ip"`
	ExpiresIn   int    `bson:"expires_in" json:"expires_in"`
}
