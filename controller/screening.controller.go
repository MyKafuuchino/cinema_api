package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type ScreeningController struct {
	screeningService service.ScreeningService
}

func NewScreeningController(screeningService service.ScreeningService) *ScreeningController {
	return &ScreeningController{screeningService: screeningService}
}

func (c *ScreeningController) GetScreenings(ctx *fiber.Ctx) error {
	screeningResponse, err := c.screeningService.GetAllScreenings()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get all screenings successfully", screeningResponse))
}

func (c *ScreeningController) GetScreeningById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}
	screeningResponse, err := c.screeningService.GetScreeningById(uId)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get screening successfully", screeningResponse))
}

func (c *ScreeningController) CreateScreening(ctx *fiber.Ctx) error {
	var req *dto.CreateScreeningRequest

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	createScreeningResponse, err := c.screeningService.CreateNewScreening(req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(types.NewResponseSuccess("Create screening successfully", createScreeningResponse))
}

func (c *ScreeningController) UpdateScreening(ctx *fiber.Ctx) error {
	var req *dto.UpdateScreeningRequest

	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)

	if err != nil {
		return err
	}
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	if err := helper.ValidateStruct(req); err != nil {
		return err
	}
	updateResponse, err := c.screeningService.UpdateScreeningById(uId, req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Update screening successfully", updateResponse))
}

func (c *ScreeningController) DeleteScreeningById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}
	deletedScreeningResponse, err := c.screeningService.DeleteScreeningById(uId)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Delete screening successfully", deletedScreeningResponse))
}

func (c *ScreeningController) GetTicketsByScreeningId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	ticketsResponse, err := c.screeningService.GetTicketsByScreeningId(uId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get tickets successfully", ticketsResponse))
}
