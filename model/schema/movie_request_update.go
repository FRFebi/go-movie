package schema

type MovieRequestUpdate struct {
	Id          int     `validate:"required" json:"id"`
	Title       string  `validate:"required" json:"title"`
	Description string  `validate:"required" json:"description"`
	Rating      float32 `validate:"required" json:"rating"`
	Image       string  `validate:"required" json:"image"`
	Update_at   string  `json:"update_at"`
}
