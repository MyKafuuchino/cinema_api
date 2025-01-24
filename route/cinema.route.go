package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/middleware"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupCinemaRoute(c fiber.Router) {
	db := database.Db
	cinemaRepo := repository.NewCinemaRepository(db)
	screeningRepo := repository.NewScreeningRepository(db)

	cinemaService := service.NewCinemaService(cinemaRepo, screeningRepo)
	cinemaController := controller.NewCinemaController(cinemaService)

	cinemaRoute := c.Group("/cinema")
	cinemaRoute.Get("", cinemaController.GetAllCinema)
	cinemaRoute.Get("/:id", cinemaController.GetCinemaById)
	cinemaRoute.Post("", middleware.ProtectRouteByRole("ADMIN"), cinemaController.CreateCinema)
	cinemaRoute.Put("/:id", middleware.ProtectRouteByRole("ADMIN"), cinemaController.UpdateCinema)
	cinemaRoute.Delete("/:id", middleware.ProtectRouteByRole("ADMIN"), cinemaController.DeleteCinemaById)

	cinemaRoute.Get("/:id/screenings", cinemaController.GetScreeningsByMovie)
}
