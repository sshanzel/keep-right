package entities

import (
	"github.com/google/uuid"
)

// Something that you have be it a container or an item
type Something struct {
	ID          uuid.UUID `pg:",type: uuid, pk"`
	Title       string
	Description string
	InsideOfID  uuid.UUID `pg:",type: uuid"`
	InsideOf    *Something
	SomeTypeID  uuid.UUID `pg:",type: uuid"`
	SomeType    *SomeType
	CreatedByID uuid.UUID `pg:",type: uuid"`
	CreatedBy   *User
}

// NewSomething creates a new instance of Something!
func NewSomething(title, description string, someTypeID, insideOfID, createdByID uuid.UUID) *Something {
	s := Something{uuid.New(), title, description, insideOfID, nil, someTypeID, nil, createdByID, nil}

	return &s
}
