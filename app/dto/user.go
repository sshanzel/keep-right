package dto

// User DTO for creating User
type User struct {
	Name string `json:"name" form:"name" query:"name"`
}
