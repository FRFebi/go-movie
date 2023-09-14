package config

import (
	"encoding/json"

	"github.com/FRFebi/go-movie/controller"
	"github.com/FRFebi/go-movie/exception"
	"github.com/FRFebi/go-movie/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(movieController controller.MovieController) *fiber.App {
	api := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Learn Restful API",
		ErrorHandler:  exception.ErrorHandler,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	movie := fiber.New()

	api.Use(recover.New())
	api.Use(middleware.Auth)
	api.Mount("api/v1", movie)
	movie.Get("/movies", movieController.FindAll)
	movie.Post("/movies", movieController.Create)
	movie.Get("/movies/:movieId", movieController.FindById)
	movie.Put("/movies/:movieId", movieController.Update)
	movie.Delete("/movies/:movieId", movieController.Delete)

	return api
}
