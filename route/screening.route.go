package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/middleware"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupScreeningRouter(r fiber.Router) {
	db := database.Db
	screeningRepo := repository.NewScreeningRepository(db)
	movieRepo := repository.NewMovieRepository(db)
	cinemaRepo := repository.NewCinemaRepository(db)
	ticketRepo := repository.NewTicketRepository(db)

	screeningService := service.NewScreeningService(screeningRepo, movieRepo, cinemaRepo, ticketRepo)
	screeningController := controller.NewScreeningController(screeningService)

	screeningRouter := r.Group("/screenings")
	screeningRouter.Get("", screeningController.GetScreenings)
	screeningRouter.Get("/:id", screeningController.GetScreeningById)
	screeningRouter.Post("", middleware.ProtectRouteByRole("ADMIN"), screeningController.CreateScreening)
	screeningRouter.Put("/:id", middleware.ProtectRouteByRole("ADMIN"), screeningController.UpdateScreening)
	screeningRouter.Delete("/:id", middleware.ProtectRouteByRole("ADMIN"), screeningController.DeleteScreeningById)

	screeningRouter.Get("/:id/tickets", screeningController.GetTicketsByScreeningId)
}
