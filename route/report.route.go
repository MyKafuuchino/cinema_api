package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupReportRouter(r fiber.Router) {
	reportRepo := repository.NewReportRepository(database.Db)
	reportService := service.NewReportService(reportRepo)
	reportController := controller.NewReportController(reportService)

	reportRoute := r.Group("/report")
	reportRoute.Post("/summary", reportController.GetSummary)
	reportRoute.Post("/cinemaSummary", reportController.GetCinemaSummary)
	reportRoute.Post("/movieSummary", reportController.GetMovieSummary)
	reportRoute.Post("/dailySummary", reportController.GetDailySummary)
}
