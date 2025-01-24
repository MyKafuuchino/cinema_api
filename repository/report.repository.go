package repository

import (
	"cinema_api/dto"
	"cinema_api/model"
	"cinema_api/types"
	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSummary(summaryReq *dto.ReportDateRequest) (*types.SummaryResponse, error)
	GetCinemaSummary(req *dto.ReportDateRequest) ([]types.CinemaSummaryResponse, error)
	GetMovieSummary(req *dto.ReportDateRequest) ([]types.MovieSummaryResponse, error)
	GetDailySummary(req *dto.ReportDateRequest) ([]types.DailySummaryResponse, error)
}

type reportRepository struct {
	db *gorm.DB
}

func (r *reportRepository) GetDailySummary(req *dto.ReportDateRequest) ([]types.DailySummaryResponse, error) {
	var dailySummaries []types.DailySummaryResponse
	err := r.db.Model(&model.Ticket{}).
		Select("DATE(tickets.created_at) AS date, COUNT(tickets.id) AS tickets_sold, SUM(screenings.price) AS total_revenue").
		Joins("JOIN screenings ON tickets.screening_id = screenings.id").
		Where("tickets.status = ?", "paid").
		Where("tickets.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Group("DATE(tickets.created_at)").
		Order("date ASC").
		Scan(&dailySummaries).Error

	return dailySummaries, err
}

func (r *reportRepository) GetMovieSummary(req *dto.ReportDateRequest) ([]types.MovieSummaryResponse, error) {
	var movieSummaries []types.MovieSummaryResponse

	err := r.db.Model(&model.Ticket{}).
		Select("movies.title AS movie_title, COUNT(tickets.id) AS tickets_sold, SUM(screenings.price) AS total_revenue, AVG(screenings.price) AS average_revenue_per_ticket").
		Joins("JOIN screenings ON tickets.screening_id = screenings.id").
		Joins("JOIN movies ON screenings.movie_id = movies.id").
		Where("tickets.status = ?", "paid").
		Where("tickets.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Group("movies.title").
		Order("tickets_sold DESC").
		Scan(&movieSummaries).Error

	return movieSummaries, err
}

func (r *reportRepository) GetCinemaSummary(req *dto.ReportDateRequest) ([]types.CinemaSummaryResponse, error) {
	var cinemaSummaries []types.CinemaSummaryResponse

	err := r.db.Model(&model.Ticket{}).
		Select("cinemas.name AS cinema_name, cinemas.location AS cinema_location, COUNT(tickets.id) AS tickets_sold, SUM(screenings.price) AS total_revenue").
		Joins("JOIN screenings ON tickets.screening_id = screenings.id").
		Joins("JOIN cinemas ON screenings.cinema_id = cinemas.id").
		Where("tickets.status = ?", "paid").
		Where("tickets.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Group("cinemas.name, cinemas.location").
		Order("tickets_sold DESC").
		Scan(&cinemaSummaries).Error

	return cinemaSummaries, err
}

func (r *reportRepository) GetSummary(summaryReq *dto.ReportDateRequest) (*types.SummaryResponse, error) {
	var summary types.SummaryResponse

	err := r.db.Model(&model.Ticket{}).
		Select("COUNT(tickets.id) AS total_tickets_sold, SUM(screenings.price) AS total_revenue").
		Joins("JOIN screenings ON tickets.screening_id = screenings.id").
		Where("tickets.status = ?", "paid").
		Where("tickets.created_at BETWEEN ? AND ?", summaryReq.StartDate, summaryReq.EndDate).
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}
