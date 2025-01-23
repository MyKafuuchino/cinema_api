package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(c fiber.Router) {
	userRepo := repository.NewUserRepository(database.Db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userRouter := c.Group("/users")
	userRouter.Get("", userController.GetAllUsers)
	userRouter.Get("/:id", userController.GetUserById)
	userRouter.Post("", userController.CreateUser)
	userRouter.Put("/:id", userController.UpdateUser)
	userRouter.Delete("/:id", userController.DeleteUserById)
}
