package model

// Session stored in the authentication service's DataBase.
type Session struct {
	ID          string `json:"id"`
	RToken      string `json:"r_token"`
	UserID      string `json:"user_id"`
	Fingerprint string `json:"fingerprint"`
	IP          string `json:"ip"`
	ExpiresIn   int    `json:"expires_in"`
}
