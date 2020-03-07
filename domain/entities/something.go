package entities

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities/something"
)

// Something that you have be it a container or an item
type Something struct {
	ID          uuid.UUID `pg:",type: uuid, pk"`
	Title       string
	Description string
	InsideOfID  uuid.UUID `pg:",type: uuid, pk"`
	InsideOf    *Something
	TypeID      uuid.UUID `pg:",type: uuid, pk"`
	Type        *something.Type
	CreatedByID uuid.UUID `pg:",type: uuid"`
	CreatedBy   *User
}

// NewSomething creates a new instance of Something!
func NewSomething(title, description string, typeID, insideOfID, createdByID uuid.UUID) *Something {
	s := Something{uuid.New(), title, description, insideOfID, nil, typeID, nil, createdByID, nil}

	return &s
}
