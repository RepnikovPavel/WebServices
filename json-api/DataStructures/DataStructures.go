package DataStructures

import "time"

type UserServerSide struct {
	UID string
}

type UserSeverPayload struct {
	NumberOfVirtualCoins uint64
	NumberOfFallenTrees  uint64
	TimeOfLastActivity   time.Time
}

type UserClientSide struct {
	FirstName string
	LastName  string
	Email     string
	TGnick    string
}

type AllUserInfo struct {
	UserSS UserServerSide
	UserCS UserClientSide
	UserSP UserSeverPayload
}

type AccountUID string

type AccountServerSide struct {
	TimeOfRegistration time.Time
}

type AllAccountInfo struct {
	AccountSS AccountServerSide
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

type ExampleStruct0 struct {
	Field0 int    `json:"innername0"`
	Field1 string `json:"innernam1"`
}

type ExampleStruct1 struct {
	Field0 string         `json:"name0"`
	Field1 int            `json:"name1"`
	Field2 string         `json:"name2"`
	Field3 float32        `json:"name3"`
	Field4 ExampleStruct0 `json:"innerStructure0"`
}

func NewExampleStruct() *ExampleStruct1 {
	return &ExampleStruct1{
		Field0: "0",
		Field1: 1,
		Field2: "2",
		Field3: 3.1,
		Field4: ExampleStruct0{
			Field0: 1,
			Field1: "1"}}
}
