package entities

import (
	"github.com/google/uuid"
)

// User information
type User struct {
	ID        uuid.UUID `sql:",type:uuid"`
	Firstname string
	Lastname  string
	Email     string
}

// NewUser return a new instance of User
func NewUser(firstname, lastname, email string) *User {
	u := User{uuid.New(), firstname, lastname, email}

	return &u
}
