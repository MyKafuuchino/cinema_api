package service

import (
	"cinema_api/dto"
	"cinema_api/repository"
	"cinema_api/types"
	"github.com/gofiber/fiber/v2"
)

type ReportService interface {
	GetSummary(reqSummary *dto.ReportDateRequest) (*types.SummaryResponse, error)
	GetCinemaSummary(req *dto.ReportDateRequest) ([]types.CinemaSummaryResponse, error)
	GetMovieSummary(req *dto.ReportDateRequest) ([]types.MovieSummaryResponse, error)
	GetDailySummary(req *dto.ReportDateRequest) ([]types.DailySummaryResponse, error)
}
type reportService struct {
	reportRepo repository.ReportRepository
}

func (s *reportService) GetDailySummary(req *dto.ReportDateRequest) ([]types.DailySummaryResponse, error) {
	dailySummaries, err := s.reportRepo.GetDailySummary(req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get Daily Summary: "+err.Error())
	}
	return dailySummaries, nil
}

func (s *reportService) GetMovieSummary(req *dto.ReportDateRequest) ([]types.MovieSummaryResponse, error) {
	movieSummaries, err := s.reportRepo.GetMovieSummary(req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get movie summary: "+err.Error())
	}
	return movieSummaries, nil
}

func (s *reportService) GetCinemaSummary(req *dto.ReportDateRequest) ([]types.CinemaSummaryResponse, error) {
	cinemaSummary, err := s.reportRepo.GetCinemaSummary(req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get cinema summary :"+err.Error())
	}
	return cinemaSummary, nil
}

func (s *reportService) GetSummary(reqSummary *dto.ReportDateRequest) (*types.SummaryResponse, error) {
	summaryResponse, err := s.reportRepo.GetSummary(reqSummary)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get summary :"+err.Error())
	}
	return summaryResponse, err
}

func NewReportService(reportRepo repository.ReportRepository) ReportService {
	return &reportService{reportRepo: reportRepo}
}
