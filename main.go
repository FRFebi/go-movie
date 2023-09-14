package main

import (
	"log"

	"github.com/FRFebi/go-movie/config"
	"github.com/FRFebi/go-movie/controller"
	"github.com/FRFebi/go-movie/repository"
	"github.com/FRFebi/go-movie/service"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	validate := validator.New()
	db := config.NewDB()

	movieRepo := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepo, db, validate)
	movieController := controller.NewMovieController(movieService)

	router := config.NewRouter(movieController)

	log.Fatal(router.Listen(":8888"))

}
