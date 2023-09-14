package repository

import (
	"context"
	"database/sql"

	"github.com/FRFebi/go-movie/model/domain"
)

type MovieRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	Delete(ctx context.Context, tx *sql.Tx, movieId int)
	Update(ctx context.Context, tx *sql.Tx, movue domain.Movie) domain.Movie
	FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.Movie, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Movie
}
