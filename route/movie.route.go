package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/middleware"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupMovieRoute(c fiber.Router) {
	movieRepo := repository.NewMovieRepository(database.Db)
	movieService := service.NewMovieService(movieRepo)
	movieController := controller.NewMovieController(movieService)

	movieRouter := c.Group("/movies")
	movieRouter.Get("", movieController.GetAllMovie)
	movieRouter.Get("/:id", movieController.GetMovieById)
	movieRouter.Post("", middleware.ProtectRouteByRole("ADMIN"), movieController.CreateMovie)
	movieRouter.Put("/:id", middleware.ProtectRouteByRole("ADMIN"), movieController.UpdateMovie)
	movieRouter.Delete("/:id", middleware.ProtectRouteByRole("ADMIN"), movieController.DeleteMovieById)
}