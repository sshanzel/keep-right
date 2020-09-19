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
}

// NewUser return a new instance of User
func NewUser(UID, name string) *User {
	u := User{uuid.New(), UID, name, false}

	return &u
}
