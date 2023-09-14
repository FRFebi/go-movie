package schema

type MovieResponse struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	Created_at  *string `json:"created_at"`
	Updated_at  *string `json:"updated_at"`
}
