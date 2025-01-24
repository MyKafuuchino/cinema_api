package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	allUserResponse, err := c.userService.GetAllUsers()
	if err != nil {
		return err
	}
	return ctx.JSON(types.NewResponseSuccess("Get all users successfully", allUserResponse))
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	userResponse, err := c.userService.GetUserById(uId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Get user by id successfully", userResponse))
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var req *dto.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	userResponse, err := c.userService.CreateUser(req)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Create user successfully", userResponse))
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	var req *dto.UpdateUserRequest

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	userResponse, err := c.userService.UpdateUserById(uId, req)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Update user successfully", userResponse))
}

func (c *UserController) DeleteUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	userResponse, err := c.userService.DeleteUserById(uId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Delete user successfully", userResponse))
}

func (c *UserController) GetTicketByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	ticketResponse, err := c.userService.GetTicketByUserId(uId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get ticket by user successfully", ticketResponse))
}
