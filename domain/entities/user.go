package entities

import (
	"github.com/google/uuid"
)

// User information
type User struct {
	ID        uuid.UUID `pg:",type:uuid, pk"`
	Firstname string
	Lastname  string
	Email     string
	IsAdmin   bool

	Somethings []Something `pg:"fk:CreatedByID"`
	SomeTypes  []SomeType  `pg:"fk:CreatedByID"`
}

// NewUser return a new instance of User
func NewUser(firstname, lastname, email string) *User {
	u := User{uuid.New(), firstname, lastname, email, false, nil, nil}

	return &u
}
