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

type MovieService interface {
	GetAllMovies() ([]types.MovieResponse, error)
	GetMovieById(id uint) (*types.MovieResponse, error)
	CreateMovie(createRequest *dto.CreateMovieRequest) (*types.MovieResponse, error)
	UpdateMovieById(id uint, updateRequest *dto.UpdateMovieRequest) (*types.MovieResponse, error)
	DeleteMovieById(id uint) (*types.MovieResponse, error)
}

type movieService struct {
	movieRepo repository.MovieRepository
}

func (s movieService) GetAllMovies() ([]types.MovieResponse, error) {
	movies, err := s.movieRepo.FindAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	allMoviesResponse := make([]types.MovieResponse, len(movies))
	for i, movie := range movies {
		allMoviesResponse[i] = types.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Genre:       movie.Genre,
			Duration:    movie.Duration,
			ReleaseDate: movie.ReleaseDate,
			CreatedAt:   movie.CreatedAt,
			UpdatedAt:   movie.UpdatedAt,
		}
	}
	return allMoviesResponse, nil
}

func (s movieService) GetMovieById(id uint) (*types.MovieResponse, error) {
	movie, err := s.movieRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Movie not found")
		}
		return nil, err
	}
	movieResponse := types.MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Genre:       movie.Genre,
		Duration:    movie.Duration,
		ReleaseDate: movie.ReleaseDate,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	return &movieResponse, nil
}

func (s movieService) CreateMovie(createRequest *dto.CreateMovieRequest) (*types.MovieResponse, error) {
	movie := &model.Movie{
		Title:       createRequest.Title,
		Description: createRequest.Description,
		Genre:       createRequest.Genre,
		Duration:    createRequest.Duration,
		ReleaseDate: createRequest.ReleaseDate,
	}

	if err := s.movieRepo.Create(movie); err != nil {
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Movie already exists : "+err.Error())
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error : "+err.Error())
	}

	movieResponse := types.MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Genre:       movie.Genre,
		Duration:    movie.Duration,
		ReleaseDate: movie.ReleaseDate,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	return &movieResponse, nil
}

func (s movieService) UpdateMovieById(id uint, updateRequest *dto.UpdateMovieRequest) (*types.MovieResponse, error) {
	movie, err := s.movieRepo.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Movie not found")
		}
		return nil, err
	}

	helper.UpdateFields(movie, updateRequest)

	if err := s.movieRepo.Update(movie); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update movie : "+err.Error())
	}

	movieResponse := types.MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Genre:       movie.Genre,
		Duration:    movie.Duration,
		ReleaseDate: movie.ReleaseDate,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	return &movieResponse, nil
}

func (s movieService) DeleteMovieById(id uint) (*types.MovieResponse, error) {
	movie, err := s.movieRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Movie not found")
		}
		return nil, err
	}

	if err := s.movieRepo.Delete(movie.ID); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete movie : "+err.Error())
	}

	movieResponse := types.MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Genre:       movie.Genre,
		Duration:    movie.Duration,
		ReleaseDate: movie.ReleaseDate,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}
	return &movieResponse, nil
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	return &movieService{movieRepo: movieRepo}
}
