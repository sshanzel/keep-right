package entities

import (
	"github.com/google/uuid"
)

// User information
type User struct {
	ID      uuid.UUID `pg:",type:uuid, pk"`
	UID     string
	Name    string
	IsAdmin bool

	Somethings []Something `pg:"fk:CreatedByID"`
	SomeTypes  []SomeType  `pg:"fk:CreatedByID"`
}

// NewUser return a new instance of User
func NewUser(UID, name string) *User {
	u := User{uuid.New(), UID, name, false, nil, nil}

	return &u
}
