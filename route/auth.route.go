package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoute(ctx fiber.Router) {
	userRepository := repository.NewUserRepository(database.Db)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	authRouter := ctx.Group("/auth")
	authRouter.Post("/login", authController.Login)
	authRouter.Post("/register", authController.Register)
}
