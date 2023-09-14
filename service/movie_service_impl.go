package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/FRFebi/go-movie/helper"
	"github.com/FRFebi/go-movie/model/domain"
	"github.com/FRFebi/go-movie/model/schema"
	"github.com/FRFebi/go-movie/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewMovieService(repository repository.MovieRepository, db *sql.DB, validate *validator.Validate) MovieService {
	return &MovieServiceImpl{
		MovieRepository: repository,
		DB:              db,
		Validate:        validate,
	}
}

func (s *MovieServiceImpl) Create(ctx context.Context, request schema.MovieRequestCreate) schema.MovieResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)
	fmt.Println("Here")

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie := domain.Movie{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
		Created_at:  StrPtr(time.Now().Format("2006-01-02- 15:04:05")),
	}

	movie = s.MovieRepository.Save(ctx, tx, movie)

	return helper.ToMovieResponse(movie)
}

func (s *MovieServiceImpl) Update(ctx context.Context, request schema.MovieRequestUpdate) schema.MovieResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := s.MovieRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}
	movie.Title = request.Title
	movie.Description = request.Description
	movie.Rating = request.Rating
	movie.Image = request.Image
	movie.Updated_at = StrPtr(time.Now().Format("2006-01-02- 15:04:05"))

	movie = s.MovieRepository.Update(ctx, tx, movie)
	return helper.ToMovieResponse(movie)
}

func (s *MovieServiceImpl) Delete(ctx context.Context, movieId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := s.MovieRepository.FindById(ctx, tx, movieId)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	s.MovieRepository.Delete(ctx, tx, movie.Id)
}

func (s *MovieServiceImpl) FindById(ctx context.Context, movieId int) schema.MovieResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := s.MovieRepository.FindById(ctx, tx, movieId)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	return helper.ToMovieResponse(movie)
}

func (s *MovieServiceImpl) FindAll(ctx context.Context) []schema.MovieResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movies := s.MovieRepository.FindAll(ctx, tx)

	return helper.ToMovieResponses(movies)
}

func StrPtr(s string) *string {
	return &s
}
