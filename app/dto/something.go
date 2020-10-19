package dto

import "github.com/google/uuid"

// Something DTO for creating Something
type Something struct {
	Title           string    `json:"title" form:"title" query:"title"`
	Description     string    `json:"description" form:"description" query:"description"`
	InsideOfID      uuid.UUID `json:"insideOfID" form:"insideOfID" query:"insideOfID"`
	SomethingTypeID uuid.UUID `json:"somethingTypeID" form:"somethingTypeID" query:"somethingTypeID"`
}
