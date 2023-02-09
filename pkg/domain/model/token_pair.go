package model

// TokenPair is a pair of both Access and Refresh tokens.
type TokenPair struct {
	AToken AccessToken
	RToken RefreshToken
}
