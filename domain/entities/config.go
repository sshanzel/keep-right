package entities

// Config Collections in Firebase
type Config struct {
	Key   string `json:"key" form:"key" query:"key"`
	Value string `json:"value" form:"value" query:"value"`
}
