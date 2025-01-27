package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/middleware"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(c fiber.Router) {
	db := database.Db
	userRepo := repository.NewUserRepository(db)
	ticketRepository := repository.NewTicketRepository(db)

	userService := service.NewUserService(userRepo, ticketRepository)
	userController := controller.NewUserController(userService)

	userRouter := c.Group("/users")
	userRouter.Get("", middleware.ProtectRouteByRole("ADMIN"), userController.GetAllUsers)
	userRouter.Get("/:id", userController.GetUserById)
	//userRouter.Post("", middleware.ProtectRouteByRole("ADMIN"), userController.CreateUser)
	userRouter.Put("/:id", middleware.ProtectRouteByRole("ADMIN"), userController.UpdateUser)
	userRouter.Delete("/:id", middleware.ProtectRouteByRole("ADMIN"), userController.DeleteUserById)

	userRouter.Get("/:id/tickets", userController.GetTicketByUserId)
}
