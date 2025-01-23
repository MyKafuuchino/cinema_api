package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req *dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	loginResponse, err := c.authService.Login(req)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("User login successfully", loginResponse))
}
