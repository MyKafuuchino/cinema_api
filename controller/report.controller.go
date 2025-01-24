package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type ReportController struct {
	reportService service.ReportService
}

func NewReportController(reportService service.ReportService) *ReportController {
	return &ReportController{reportService: reportService}
}

func (c *ReportController) GetSummary(ctx *fiber.Ctx) error {
	var req *dto.ReportDateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body :"+err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	summaryResponse, err := c.reportService.GetSummary(req)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get summary successfully", summaryResponse))
}

func (c *ReportController) GetCinemaSummary(ctx *fiber.Ctx) error {
	var req *dto.ReportDateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body :"+err.Error())
	}
	if err := helper.ValidateStruct(req); err != nil {
		return err
	}
	cinemaSummaryResponse, err := c.reportService.GetCinemaSummary(req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get cinema summary successfully", cinemaSummaryResponse))
}

func (c *ReportController) GetMovieSummary(ctx *fiber.Ctx) error {
	var req *dto.ReportDateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body :"+err.Error())
	}
	if err := helper.ValidateStruct(req); err != nil {
		return err
	}
	movieSummaryResponse, err := c.reportService.GetMovieSummary(req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get movie summary successfully", movieSummaryResponse))
}

func (c *ReportController) GetDailySummary(ctx *fiber.Ctx) error {
	var req *dto.ReportDateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body :"+err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	dailySummaryResponse, err := c.reportService.GetDailySummary(req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get daily summary successfully", dailySummaryResponse))
}
