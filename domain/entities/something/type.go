package something

import "github.com/google/uuid"

// Type is the struct for identifying a type of Something
type Type struct {
	ID          uuid.UUID `pg:",type: uuid"`
	Title       string
	Description string
}
