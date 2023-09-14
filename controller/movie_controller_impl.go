package controller

import (
	"strconv"

	"github.com/FRFebi/go-movie/helper"
	"github.com/FRFebi/go-movie/model/schema"
	"github.com/FRFebi/go-movie/service"
	"github.com/gofiber/fiber/v2"
)

type MovieControllerImpl struct {
	MovieService service.MovieService
}

func NewMovieController(service service.MovieService) MovieController {
	return &MovieControllerImpl{
		MovieService: service,
	}
}

func (c *MovieControllerImpl) Create(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	movieRequest := schema.MovieRequestCreate{}
	err := ctx.BodyParser(&movieRequest)
	helper.PanicIfError(err)

	movie := c.MovieService.Create(ctx.UserContext(), movieRequest)
	response := schema.ApiResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    movie,
	}

	return helper.WriteResponseBody(ctx, response)
}

func (c *MovieControllerImpl) Update(ctx *fiber.Ctx) error {
	movieId, errParseInt := strconv.Atoi(ctx.Params("movieId"))
	helper.PanicIfError(errParseInt)

	movieRequest := schema.MovieRequestUpdate{}
	err := ctx.BodyParser(&movieRequest)
	helper.PanicIfError(err)

	movieRequest.Id = movieId
	movie := c.MovieService.Update(ctx.UserContext(), movieRequest)
	response := schema.ApiResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    movie,
	}

	return helper.WriteResponseBody(ctx, response)
}

func (c *MovieControllerImpl) Delete(ctx *fiber.Ctx) error {
	movieId, errParseInt := strconv.Atoi(ctx.Params("movieId"))
	helper.PanicIfError(errParseInt)

	c.MovieService.Delete(ctx.UserContext(), movieId)
	response := schema.ApiResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
	}

	return helper.WriteResponseBody(ctx, response)
}

func (c *MovieControllerImpl) FindById(ctx *fiber.Ctx) error {
	movieId, errParseInt := strconv.Atoi(ctx.Params("movieId"))
	helper.PanicIfError(errParseInt)
	movie := c.MovieService.FindById(ctx.UserContext(), movieId)
	response := schema.ApiResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    movie,
	}

	return helper.WriteResponseBody(ctx, response)
}

func (c *MovieControllerImpl) FindAll(ctx *fiber.Ctx) error {
	movie := c.MovieService.FindAll(ctx.UserContext())
	response := schema.ApiResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    movie,
	}

	return helper.WriteResponseBody(ctx, response)
}
