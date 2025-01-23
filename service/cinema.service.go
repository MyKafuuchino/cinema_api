package service

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/model"
	"cinema_api/repository"
	"cinema_api/types"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CinemaService interface {
	GetAllCinema() ([]types.CinemaResponse, error)
	GetCinemaById(id uint) (*types.CinemaResponse, error)
	CreateCinema(createRequest *dto.CreateCinemaRequest) (*types.CinemaResponse, error)
	UpdateCinemaById(id uint, updateRequest *dto.UpdateCinemaRequest) (*types.CinemaResponse, error)
	DeleteCinemaById(id uint) (*types.CinemaResponse, error)
}

type cinemaService struct {
	cinemaRepo repository.CinemaRepository
}

func (s *cinemaService) GetAllCinema() ([]types.CinemaResponse, error) {
	cinemas, err := s.cinemaRepo.FindAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	allCinemasResponse := make([]types.CinemaResponse, len(cinemas))
	for i, cinema := range cinemas {
		allCinemasResponse[i] = types.CinemaResponse{
			ID:        cinema.ID,
			Name:      cinema.Name,
			Location:  cinema.Location,
			CreatedAt: cinema.CreatedAt,
			UpdatedAt: cinema.UpdatedAt,
		}
	}
	return allCinemasResponse, nil
}

func (s *cinemaService) GetCinemaById(id uint) (*types.CinemaResponse, error) {
	cinema, err := s.cinemaRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Cinema not found")
		}
		return nil, err
	}
	cinemaResponse := types.CinemaResponse{
		ID:        cinema.ID,
		Name:      cinema.Name,
		Location:  cinema.Location,
		CreatedAt: cinema.CreatedAt,
		UpdatedAt: cinema.UpdatedAt,
	}

	return &cinemaResponse, nil
}

func (s *cinemaService) CreateCinema(createRequest *dto.CreateCinemaRequest) (*types.CinemaResponse, error) {
	cinema := &model.Cinema{
		Name:     createRequest.Name,
		Location: createRequest.Location,
	}

	if err := s.cinemaRepo.Create(cinema); err != nil {
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Cinema already exists : "+err.Error())
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error : "+err.Error())
	}

	cinemaResponse := types.CinemaResponse{
		ID:        cinema.ID,
		Name:      cinema.Name,
		Location:  cinema.Location,
		CreatedAt: cinema.CreatedAt,
		UpdatedAt: cinema.UpdatedAt,
	}

	return &cinemaResponse, nil
}

func (s *cinemaService) UpdateCinemaById(id uint, updateRequest *dto.UpdateCinemaRequest) (*types.CinemaResponse, error) {
	cinema, err := s.cinemaRepo.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Cinema not found")
		}
		return nil, err
	}

	helper.UpdateFields(cinema, updateRequest)

	if err := s.cinemaRepo.Update(cinema); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update cinema: "+err.Error())
	}

	cinemaResponse := types.CinemaResponse{
		ID:        cinema.ID,
		Name:      cinema.Name,
		Location:  cinema.Location,
		CreatedAt: cinema.CreatedAt,
		UpdatedAt: cinema.UpdatedAt,
	}

	return &cinemaResponse, nil
}

func (s *cinemaService) DeleteCinemaById(id uint) (*types.CinemaResponse, error) {
	cinema, err := s.cinemaRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Cinema not found")
		}
		return nil, err
	}

	if err := s.cinemaRepo.Delete(cinema.ID); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete cinema : "+err.Error())
	}

	cinemaResponse := types.CinemaResponse{
		ID:        cinema.ID,
		Name:      cinema.Name,
		Location:  cinema.Location,
		CreatedAt: cinema.CreatedAt,
		UpdatedAt: cinema.UpdatedAt,
	}
	return &cinemaResponse, nil
}

func NewCinemaService(cinemaRepo repository.CinemaRepository) CinemaService {
	return &cinemaService{cinemaRepo: cinemaRepo}
}
