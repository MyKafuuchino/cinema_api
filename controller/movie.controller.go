package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type MovieController struct {
	movieService service.MovieService
}

func NewMovieController(movieService service.MovieService) *MovieController {
	return &MovieController{movieService: movieService}
}

func (c *MovieController) GetAllMovie(ctx *fiber.Ctx) error {
	allMovieResponse, err := c.movieService.GetAllMovies()
	if err != nil {
		return err
	}
	return ctx.JSON(types.NewResponseSuccess("Get all movie successfully", allMovieResponse))
}

func (c *MovieController) GetMovieById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	movieResponse, err := c.movieService.GetMovieById(uId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Get movie by id successfully", movieResponse))
}

func (c *MovieController) CreateMovie(ctx *fiber.Ctx) error {
	var req *dto.CreateMovieRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	movieResponse, err := c.movieService.CreateMovie(req)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Create movie successfully", movieResponse))
}

func (c *MovieController) UpdateMovie(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	var req *dto.UpdateMovieRequest

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	movieResponse, err := c.movieService.UpdateMovieById(uId, req)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Update movie successfully", movieResponse))
}

func (c *MovieController) DeleteMovieById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	movieResponse, err := c.movieService.DeleteMovieById(uId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.NewResponseSuccess("Delete movie successfully", movieResponse))
}
