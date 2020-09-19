package entities

import (
	"github.com/google/uuid"
)

// Something that you have be it a container or an item
type Something struct {
	ID              uuid.UUID `pg:",type: uuid, pk"`
	Title           string
	Description     string
	InsideOfID      uuid.UUID `pg:",type: uuid"`
	InsideOf        *Something
	SomethingTypeID uuid.UUID `pg:",type: uuid"`
	SomethingType   *SomethingType
	CreatedByID     uuid.UUID `pg:",type: uuid"`
	CreatedBy       *User
}

// NewSomething creates a new instance of Something!
func NewSomething(title, description string, somethingTypeID, insideOfID, createdByID uuid.UUID) *Something {
	s := Something{uuid.New(), title, description, insideOfID, nil, somethingTypeID, nil, createdByID, nil}

	return &s
}
