package controller

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/service"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TicketController struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (c *TicketController) CreateTicket(ctx *fiber.Ctx) error {
	var req *dto.CreateTicketRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse body: "+err.Error())
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
	ticketsResponse, err := c.ticketService.GetAllTickets()
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
