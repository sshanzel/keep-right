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
}

// NewUser return a new instance of User
func NewUser(firstname, lastname, email string) *User {
	u := User{uuid.New(), firstname, lastname, email, false}

	return &u
}
