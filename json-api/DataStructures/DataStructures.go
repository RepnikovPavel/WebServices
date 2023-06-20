package DataStructures

import "time"

type UserAccount struct {
	CurrentUserName   string
	CurrentUserSecret string
	AccountID         string
}

type Event_CreateAccount struct {
	Time       time.Time
	UserName   string
	UserSecret string
}

type Event_DeleteAccount struct {
	Time       time.Time
	UserName   string
	UserSecret string
}
