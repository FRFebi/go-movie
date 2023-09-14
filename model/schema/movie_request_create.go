package schema

type MovieRequestCreate struct {
	Title       string  `validate:"required" json:"title"`
	Description string  `validate:"required" json:"description"`
	Rating      float32 `validate:"required" json:"rating"`
	Image       string  `validate:"required" json:"image"`
	Created_at  string  `json:"created_at"`
}
