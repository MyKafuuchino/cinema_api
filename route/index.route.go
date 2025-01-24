package route

import (
	"github.com/gofiber/fiber/v2"
)

func InitRouter(router fiber.Router) {
	api := router.Group("/api")
	SetupAuthRoute(api)
	SetupUserRouter(api)
	SetupMovieRoute(api)
	SetupCinemaRoute(api)
	SetupScreeningRouter(api)
}
