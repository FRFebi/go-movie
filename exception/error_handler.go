package exception

import (
	"errors"

	"github.com/FRFebi/go-movie/model/schema"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if e, ok := err.(*fiber.Error); ok {
		if e.Code == fiber.ErrNotFound.Code {
			return notFoundError(ctx, err)
		}
	}
	if err, ok := err.(validator.ValidationErrors); ok {
		return validationError(ctx, err)
	}

	return internalServerError(ctx, err)
}

func validationError(ctx *fiber.Ctx, err error) error {
	newError := fiber.NewError(fiber.StatusBadRequest)
	code := newError.Code
	status := newError.Message
	webResponse := schema.ApiResponse{
		Code:    code,
		Message: status,
		Data:    err.Error(),
	}

	err = ctx.Status(code).JSON(webResponse)
	if err != nil {
		internalServerError(ctx, err)
	}
	return nil
}

func internalServerError(ctx *fiber.Ctx, err error) error {
	fError := fiber.ErrInternalServerError
	code := fError.Code
	status := fError.Message
	var data interface{} = nil

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		data = e.Message
	}

	webResponse := schema.ApiResponse{
		Code:    code,
		Message: status,
		Data:    data,
	}

	err = ctx.Status(code).JSON(webResponse)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("INTERNAL SERVER ERROR")
	}
	return nil
}

func notFoundError(ctx *fiber.Ctx, err error) error {
	ferror := fiber.ErrNotFound
	code := ferror.Code
	var data interface{} = nil

	e, ok := err.(*fiber.Error)
	if ok {
		code = e.Code
		data = e.Message
	}

	webResponse := schema.ApiResponse{
		Code:    code,
		Message: fiber.ErrNotFound.Message,
		Data:    data,
	}

	err = ctx.Status(code).JSON(webResponse)
	if err != nil {
		return internalServerError(ctx, err)
	}
	return nil
}
