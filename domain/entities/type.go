package entities

import (
	"github.com/google/uuid"
)

// SomethingType is the struct for identifying a type of Something
type SomethingType struct {
	ID          uuid.UUID `pg:",type: uuid, pk"`
	Title       string
	Description string
	CreatedByID uuid.UUID `pg:",type: uuid"`
	CreatedBy   *User
}

// NewSomeType creates a new Type of Something
func NewSomeType(title, description string, createdByID uuid.UUID) *SomethingType {
	t := SomethingType{uuid.New(), title, description, createdByID, nil}

	return &t
}
