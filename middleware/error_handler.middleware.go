package middleware

import (
	"cinema_api/helper"
	"cinema_api/types"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorList := helper.GeValidationErrors(validationErrors)
			return ctx.Status(fiber.StatusBadRequest).JSON(types.NewResponseError("Validation Error", errorList...))
		}

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			return ctx.Status(code).JSON(types.NewResponseError(fiberErr.Message))
		}

		return ctx.Status(code).JSON(types.NewResponseError("Internal Server Error : ", err.Error()))
	}
}
