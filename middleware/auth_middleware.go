package middleware

import (
	"encoding/base64"

	"github.com/FRFebi/go-movie/model/schema"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	if ctx.Get("Authorization") == "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")) {
		return ctx.Next()
	}

	code := fiber.ErrUnauthorized.Code
	status := fiber.ErrUnauthorized.Message
	webResponse := schema.ApiResponse{
		Code:    code,
		Message: status,
		Data:    nil,
	}

	err := ctx.Status(code).JSON(webResponse)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("INTERNAL SERVER ERROR")
	}
	return nil
}
