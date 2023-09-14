package service

import (
	"context"

	"github.com/FRFebi/go-movie/model/schema"
)

type MovieService interface {
	Create(ctx context.Context, request schema.MovieRequestCreate) schema.MovieResponse
	Update(ctx context.Context, request schema.MovieRequestUpdate) schema.MovieResponse
	Delete(ctx context.Context, movieId int)
	FindById(ctx context.Context, movieId int) schema.MovieResponse
	FindAll(ctx context.Context) []schema.MovieResponse
}
