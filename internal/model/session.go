package model

type Session struct {
	RToken      RefreshToken
	UserId      userID
	FingerPrint string
	Ip          string
	ExperisIn   int
}
