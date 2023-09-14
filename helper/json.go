package helper

import (
	"github.com/FRFebi/go-movie/model/schema"
	"github.com/gofiber/fiber/v2"
)

func WriteResponseBody(c *fiber.Ctx, response schema.ApiResponse) error {
	c.Set("Content-Type", "application/json")
	return c.JSON(response)
}
