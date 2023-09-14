package domain

type Movie struct {
	Id          int
	Title       string
	Description string
	Rating      float32
	Image       string
	Created_at  *string
	Updated_at  *string
}
