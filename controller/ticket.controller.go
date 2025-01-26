package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type TicketController struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (c *TicketController) CreateTicket(ctx *fiber.Ctx) error {
	claim := ctx.Locals("claim")
	user := claim.(types.UserPayload)

	var req *dto.CreateTicketRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body: "+err.Error())
	}

	if user.Id != req.UserID {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user")
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	createResponse, err := c.ticketService.Create(req)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(types.NewResponseSuccess("Create Ticket Successfully", createResponse))
}

func (c *TicketController) GetAllTickets(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, "Limit parameter must be an integer")
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, "Offset parameter must be an integer")
	}
	params := &types.QueryParamRequest{
		Status: ctx.Query("status"),
		Limit:  limit,
		Offset: offset,
	}
	ticketsResponse, err := c.ticketService.GetAllTickets(params)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(types.NewResponseSuccess("Get Tickets Successfully", ticketsResponse))
}

func (c *TicketController) GetTicketById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	ticketResponse, err := c.ticketService.GetTicketById(uId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Get ticket by id Successfully", ticketResponse))
}

func (c *TicketController) UpdateTicketById(ctx *fiber.Ctx) error {
	var req *dto.UpdateTicketRequest

	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body: "+err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	updateResponse, err := c.ticketService.UpdateTicketById(uId, req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Update Ticket Successfully", updateResponse))
}

func (c *TicketController) DeleteTicket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	deleteResponse, err := c.ticketService.DeleteTicketById(uId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Ticket Delete Successfully", deleteResponse))
}

func (c *TicketController) UpdateTicketStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uId, err := helper.StringToUint(id)
	if err != nil {
		return err
	}

	status := ctx.Query("status")

	updateResponse, err := c.ticketService.UpdateTicketStatus(uId, status)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(types.NewResponseSuccess("Ticket Update Successfully", updateResponse))
}
