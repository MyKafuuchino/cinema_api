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

type ScreeningService interface {
	GetAllScreenings() ([]types.ScreeningResponse, error)
	GetScreeningById(id uint) (*types.ScreeningResponse, error)
	CreateNewScreening(screeningReq *dto.CreateScreeningRequest) (*types.ScreeningResponse, error)
	UpdateScreeningById(id uint, screeningReq *dto.UpdateScreeningRequest) (*types.ScreeningResponse, error)
	DeleteScreeningById(id uint) (*types.ScreeningResponse, error)

	GetTicketsByScreeningId(id uint) ([]types.TicketResponse, error)
}

type screeningService struct {
	screeningRepo repository.ScreeningRepository
	movieRepo     repository.MovieRepository
	cinemaRepo    repository.CinemaRepository
	ticketRepo    repository.TicketRepository
}

func (s *screeningService) GetAllScreenings() ([]types.ScreeningResponse, error) {
	screenings, err := s.screeningRepo.FindAll()
	if err != nil {
		return nil, err
	}
	screeningsResponse := make([]types.ScreeningResponse, len(screenings))
	for i, screening := range screenings {
		screeningsResponse[i] = types.ScreeningResponse{
			ID:            screening.ID,
			MovieID:       screening.MovieID,
			CinemaID:      screening.CinemaID,
			ScreeningTime: screening.ScreeningTime,
			Price:         screening.Price,
			CreatedAt:     screening.CreatedAt,
			UpdatedAt:     screening.UpdatedAt,
		}
	}
	return screeningsResponse, nil
}

func (s *screeningService) GetScreeningById(id uint) (*types.ScreeningResponse, error) {
	screening, err := s.screeningRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "screening not found : "+err.Error())
		}
		return nil, err
	}

	screeningResponse := types.ScreeningResponse{
		ID:            screening.ID,
		MovieID:       screening.MovieID,
		CinemaID:      screening.CinemaID,
		ScreeningTime: screening.ScreeningTime,
		Price:         screening.Price,
		CreatedAt:     screening.CreatedAt,
		UpdatedAt:     screening.UpdatedAt,
	}

	return &screeningResponse, nil
}

func (s *screeningService) CreateNewScreening(screeningReq *dto.CreateScreeningRequest) (*types.ScreeningResponse, error) {
	movie, err := s.movieRepo.FindById(screeningReq.MovieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "movie not found : "+err.Error())
		}
		return nil, err
	}
	cinema, err := s.cinemaRepo.FindById(screeningReq.CinemaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "cinema not found : "+err.Error())
		}
		return nil, err
	}

	screening := model.Screening{
		MovieID:       movie.ID,
		CinemaID:      cinema.ID,
		ScreeningTime: screeningReq.ScreeningTime,
		Price:         screeningReq.Price,
	}

	if err := s.screeningRepo.Create(&screening); err != nil {
		return nil, err
	}

	screeningResponse := types.ScreeningResponse{
		ID:            screening.ID,
		MovieID:       screening.MovieID,
		CinemaID:      screening.CinemaID,
		ScreeningTime: screening.ScreeningTime,
		Price:         screening.Price,
		CreatedAt:     screening.CreatedAt,
		UpdatedAt:     screening.UpdatedAt,
	}

	return &screeningResponse, nil
}

func (s *screeningService) UpdateScreeningById(id uint, screeningReq *dto.UpdateScreeningRequest) (*types.ScreeningResponse, error) {
	screening, err := s.screeningRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "screening not found", err.Error())
		}
		return nil, err
	}

	if screeningReq.MovieID != nil {
		screening.MovieID = *screeningReq.MovieID
	}

	if screeningReq.CinemaID != nil {
		screening.CinemaID = *screeningReq.CinemaID
	}

	if screeningReq.ScreeningTime != nil {
		screening.ScreeningTime = *screeningReq.ScreeningTime
	}

	if screeningReq.Price != nil {
		screening.Price = *screeningReq.Price
	}

	_, err = s.movieRepo.FindById(screening.MovieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "movie not found : "+err.Error())
		}
		return nil, err
	}

	_, err = s.cinemaRepo.FindById(screening.CinemaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "cinema not found : "+err.Error())
		}
		return nil, err
	}

	if err := s.screeningRepo.Update(screening); err != nil {
		return nil, err
	}

	screeningResponse := types.ScreeningResponse{
		ID:            screening.ID,
		MovieID:       screening.MovieID,
		CinemaID:      screening.CinemaID,
		ScreeningTime: screening.ScreeningTime,
		Price:         screening.Price,
		CreatedAt:     screening.CreatedAt,
		UpdatedAt:     screening.UpdatedAt,
	}

	return &screeningResponse, nil
}

func (s *screeningService) DeleteScreeningById(id uint) (*types.ScreeningResponse, error) {
	screening, err := s.screeningRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "screening not found : "+err.Error())
		}
		return nil, err
	}

	if err := s.screeningRepo.Delete(screening.ID); err != nil {
		return nil, err
	}

	screeningResponse := types.ScreeningResponse{
		ID:            screening.ID,
		MovieID:       screening.MovieID,
		CinemaID:      screening.CinemaID,
		ScreeningTime: screening.ScreeningTime,
		Price:         screening.Price,
		CreatedAt:     screening.CreatedAt,
		UpdatedAt:     screening.UpdatedAt,
	}
	return &screeningResponse, nil
}

func (s *screeningService) GetTicketsByScreeningId(id uint) ([]types.TicketResponse, error) {
	if _, err := s.screeningRepo.FindById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "screening not found : "+err.Error())
		}
		return nil, err
	}

	tickets, err := s.ticketRepo.FindByScreeningId(id)
	if err != nil {
		return nil, err
	}

	ticketsResponse := make([]types.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		ticketsResponse[i] = types.TicketResponse{
			ID:          ticket.ID,
			UserID:      ticket.UserID,
			ScreeningID: ticket.ScreeningID,
			SeatNumber:  ticket.SeatNumber,
			Status:      ticket.Status,
			CreatedAt:   ticket.CreatedAt,
			UpdatedAt:   ticket.UpdatedAt,
		}
	}
	return ticketsResponse, nil
}

func NewScreeningService(screeningRepo repository.ScreeningRepository, movieRepo repository.MovieRepository, cinemaRepo repository.CinemaRepository, ticketRepository repository.TicketRepository) ScreeningService {
	return &screeningService{screeningRepo: screeningRepo, movieRepo: movieRepo, cinemaRepo: cinemaRepo, ticketRepo: ticketRepository}
}
