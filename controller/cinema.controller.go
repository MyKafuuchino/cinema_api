package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type CinemaController struct {
	cinemaService service.CinemaService
}

func NewCinemaController(cinemaService service.CinemaService) *CinemaController {
	return &CinemaController{cinemaService: cinemaService}
}

func (c *CinemaController) GetAllCinema(ctx *fiber.Ctx) error {
	allCinemaResponse, err := c.cinemaService.GetAllCinema()
	if err != nil {
		return err
	}
	return ctx.JSON(types.NewResponseSuccess("Get all cinema successfully", allCinemaResponse))
}

func (c *CinemaController) GetCinemaById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	cinemaResponse, err := c.cinemaService.GetCinemaById(uId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Get cinema by id successfully", cinemaResponse))
}

func (c *CinemaController) CreateCinema(ctx *fiber.Ctx) error {
	var req *dto.CreateCinemaRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	cinemaResponse, err := c.cinemaService.CreateCinema(req)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Create cinema successfully", cinemaResponse))
}

func (c *CinemaController) UpdateCinema(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	var req *dto.UpdateCinemaRequest

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	cinemaResponse, err := c.cinemaService.UpdateCinemaById(uId, req)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Update cinema successfully", cinemaResponse))
}

func (c *CinemaController) DeleteCinemaById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	cinemaResponse, err := c.cinemaService.DeleteCinemaById(uId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Delete cinema successfully", cinemaResponse))
}
