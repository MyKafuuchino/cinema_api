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

type MovieService interface {
	GetAllMovies() ([]types.MovieResponse, error)
	GetMovieById(id uint) (*types.MovieResponse, error)
	CreateMovie(createRequest *dto.CreateMovieRequest) (*types.MovieResponse, error)
	UpdateMovieById(id uint, updateRequest *dto.UpdateMovieRequest) (*types.MovieResponse, error)
	DeleteMovieById(id uint) (*types.MovieResponse, error)

	GetScreeningsByMovie(id uint) ([]types.ScreeningResponse, error)
}

type movieService struct {
	movieRepo     repository.MovieRepository
	screeningRepo repository.ScreeningRepository
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

	if updateRequest.Title != nil {
		movie.Title = *updateRequest.Title
	}
	if updateRequest.Description != nil {
		movie.Description = *updateRequest.Description
	}
	if updateRequest.Genre != nil {
		movie.Genre = *updateRequest.Genre
	}
	if updateRequest.Duration != nil {
		movie.Duration = *updateRequest.Duration
	}
	if updateRequest.ReleaseDate != nil {
		movie.ReleaseDate = *updateRequest.ReleaseDate
	}

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

func (s movieService) GetScreeningsByMovie(id uint) ([]types.ScreeningResponse, error) {
	screenings, err := s.screeningRepo.FindByCinema(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Screening not found")
		}
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

func NewMovieService(movieRepo repository.MovieRepository, screeningRepo repository.ScreeningRepository) MovieService {
	return &movieService{movieRepo: movieRepo, screeningRepo: screeningRepo}
}
