package entities

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities/something"
)

// Something that you have be it a container or an item
type Something struct {
	ID          uuid.UUID `pg:",type: uuid"`
	Title       string
	Description string
	InsideOfID  uuid.UUID
	InsideOf    *Something
	TypeID      uuid.UUID
	Type        *something.Type
}

// NewSomething creates a new instance of Something!
func NewSomething(title, description string, typeID, insideOfID uuid.UUID) *Something {
	s := Something{uuid.New(), title, description, insideOfID, nil, typeID, nil}

	return &s
}
