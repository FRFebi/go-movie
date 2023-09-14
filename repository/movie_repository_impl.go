package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FRFebi/go-movie/helper"
	"github.com/FRFebi/go-movie/model/domain"
)

type MovieRepositoryImpl struct{}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}

func (r *MovieRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	sql := "INSERT INTO movie (title, description, rating, image,created_at) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, movie.Title, movie.Description, movie.Rating, movie.Image, movie.Created_at)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	movie.Id = int(id)

	return movie
}

func (r *MovieRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	sql := "UPDATE movie SET title=?, description=?, rating=?, image=?, updated_at=? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, movie.Title, movie.Description, movie.Rating, movie.Image, movie.Updated_at, movie.Id)
	helper.PanicIfError(err)

	return movie
}

func (r *MovieRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, movieId int) {
	sql := "DELETE FROM movie where id=?"
	_, err := tx.ExecContext(ctx, sql, movieId)
	helper.PanicIfError(err)
}

func (r *MovieRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.Movie, error) {
	sql := "SELECT id, title, description, rating, image, created_at, updated_at FROM movie WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, movieId)
	helper.PanicIfError(err)
	defer rows.Close()

	movie := domain.Movie{}
	if rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &movie.Created_at, &movie.Updated_at)
		helper.PanicIfError(err)
		return movie, err
	}

	return movie, errors.New("movie not found")
}

func (r *MovieRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Movie {
	sql := "SELECT id, title, description, rating, image, created_at, updated_at FROM movie LIMIT 5"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	movies := []domain.Movie{}
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &movie.Created_at, &movie.Updated_at)
		helper.PanicIfError(err)
		movies = append(movies, movie)
	}
	return movies
}
