package helper

import (
	"github.com/FRFebi/go-movie/model/domain"
	"github.com/FRFebi/go-movie/model/schema"
)

func ToMovieResponse(movie domain.Movie) schema.MovieResponse {
	return schema.MovieResponse{
		Id:          movie.Id,
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
		Created_at:  movie.Created_at,
		Updated_at:  movie.Updated_at,
	}
}

func ToMovieResponses(movies []domain.Movie) []schema.MovieResponse {
	moviesResponses := []schema.MovieResponse{}
	for _, v := range movies {
		moviesResponses = append(moviesResponses, ToMovieResponse(v))
	}
	return moviesResponses
}
