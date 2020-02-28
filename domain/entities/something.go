package entities

import (
	"github.com/google/uuid"
)

// Something that you have be it a container or an item
type Something struct {
	ID          uuid.UUID
	Title       string
	Description string
}

// NewSomething creates a new instance of Something!
func NewSomething(title, description string) *Something {
	s := Something{uuid.New(), title, description}

	return &s
}
