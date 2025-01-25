package service

import (
	"cinema_api/dto"
	"cinema_api/model"
	"cinema_api/repository"
	"cinema_api/types"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TicketService interface {
	Create(ticketReq *dto.CreateTicketRequest) (*types.TicketResponse, error)
	GetAllTickets(params *types.QueryParamRequest) ([]types.TicketResponse, error)
	GetTicketById(id uint) (*types.TicketResponse, error)
	UpdateTicketById(id uint, ticketReq *dto.UpdateTicketRequest) (*types.TicketResponse, error)
	DeleteTicketById(id uint) (*types.TicketResponse, error)

	UpdateTicketStatus(id uint, status string) (*types.TicketResponse, error)
}

type ticketService struct {
	ticketRepo    repository.TicketRepository
	userRepo      repository.UserRepository
	screeningRepo repository.ScreeningRepository
}

func (s *ticketService) UpdateTicketStatus(id uint, status string) (*types.TicketResponse, error) {
	isTicket, err := s.ticketRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Ticket not found :"+err.Error())
		}
		return nil, err
	}

	switch status {
	case "paid":
		isTicket.Status = status
	case "booked":
		isTicket.Status = status
	case "canceled":
		isTicket.Status = status
	default:
		return nil, fiber.NewError(fiber.StatusBadRequest, "Error updating ticket status :"+status)
	}

	if err := s.ticketRepo.Update(isTicket); err != nil {
		return nil, err
	}

	ticketResponse := &types.TicketResponse{
		ID:          isTicket.ID,
		UserID:      isTicket.UserID,
		ScreeningID: isTicket.ScreeningID,
		SeatNumber:  isTicket.SeatNumber,
		Status:      isTicket.Status,
		CreatedAt:   isTicket.CreatedAt,
		UpdatedAt:   isTicket.UpdatedAt,
	}

	return ticketResponse, nil
}

func (s *ticketService) Create(ticketReq *dto.CreateTicketRequest) (*types.TicketResponse, error) {
	user, err := s.userRepo.FindById(ticketReq.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return nil, err
	}

	screening, err := s.screeningRepo.FindById(ticketReq.ScreeningID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Screening not found")
		}
		return nil, err
	}

	ticket := &model.Ticket{
		UserID:      user.ID,
		ScreeningID: screening.ID,
		SeatNumber:  ticketReq.SeatNumber,
		Status:      ticketReq.Status,
	}

	if err := s.ticketRepo.Create(ticket); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusConflict, "Seat already booked")
		}
		return nil, err
	}

	createResponse := &types.TicketResponse{
		ID:          ticket.ID,
		UserID:      ticket.UserID,
		ScreeningID: ticket.ScreeningID,
		SeatNumber:  ticket.SeatNumber,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}

	return createResponse, nil
}

func (s *ticketService) GetAllTickets(params *types.QueryParamRequest) ([]types.TicketResponse, error) {
	ticket, err := s.ticketRepo.FindAll(params)
	if err != nil {
		return nil, err
	}
	ticketResponse := make([]types.TicketResponse, len(ticket))
	for i, item := range ticket {
		ticketResponse[i] = types.TicketResponse{
			ID:          item.ID,
			UserID:      item.UserID,
			ScreeningID: item.ScreeningID,
			SeatNumber:  item.SeatNumber,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}
	return ticketResponse, nil
}

func (s *ticketService) GetTicketById(id uint) (*types.TicketResponse, error) {
	ticket, err := s.ticketRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Ticket not found")
		}
		return nil, err
	}
	ticketResponse := &types.TicketResponse{
		ID:          ticket.ID,
		UserID:      ticket.UserID,
		ScreeningID: ticket.ScreeningID,
		SeatNumber:  ticket.SeatNumber,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}
	return ticketResponse, nil
}

func (s *ticketService) UpdateTicketById(id uint, ticketReq *dto.UpdateTicketRequest) (*types.TicketResponse, error) {
	ticket, err := s.ticketRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Ticket not found")
		}
		return nil, err
	}

	if ticketReq != nil {
		ticket.SeatNumber = *ticketReq.SeatNumber
	}

	if err := s.ticketRepo.Update(ticket); err != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Ticket update failed :"+err.Error())
	}

	updateResponse := &types.TicketResponse{
		ID:          ticket.ID,
		UserID:      ticket.UserID,
		ScreeningID: ticket.ScreeningID,
		SeatNumber:  ticket.SeatNumber,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}

	return updateResponse, nil
}

func (s *ticketService) DeleteTicketById(id uint) (*types.TicketResponse, error) {
	ticket, err := s.ticketRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Ticket not found")
		}
		return nil, err
	}
	if err := s.ticketRepo.Delete(ticket); err != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Ticket delete failed :"+err.Error())
	}
	ticketResponse := &types.TicketResponse{
		ID:          ticket.ID,
		UserID:      ticket.UserID,
		ScreeningID: ticket.ScreeningID,
		SeatNumber:  ticket.SeatNumber,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}
	return ticketResponse, nil
}

func NewTicketService(ticketRepo repository.TicketRepository, userRepo repository.UserRepository, screeningRepo repository.ScreeningRepository) TicketService {
	return &ticketService{ticketRepo: ticketRepo, userRepo: userRepo, screeningRepo: screeningRepo}
}
