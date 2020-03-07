package something

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities"
)

// Type is the struct for identifying a type of Something
type Type struct {
	ID          uuid.UUID `pg:",type: uuid, pk"`
	Title       string
	Description string
	CreatedByID uuid.UUID `pg:",type: uuid"`
	CreatedBy   *entities.User
}

// NewType creates a new Type of Something
func NewType(title, description string, createdByID uuid.UUID) *Type {
	t := Type{uuid.New(), title, description, createdByID, nil}

	return &t
}
