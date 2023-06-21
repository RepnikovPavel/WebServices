package DataStructures

import "time"

type UserServerSide struct {
	UID string
}
type UserSeverPayload struct {
	NumberOfVirtulCoins uint64
	NumberOfFallenTrees uint64
	TimeOfLastActivity  time.Time
}

type UserClientSide struct {
	FirstName string
	LastName  string
	Email     string
	TGnick    string
}

type UserAccount struct {
	UserUID   string
	AccountID string
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
