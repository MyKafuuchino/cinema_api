package route

import (
	userRouter "cinema_api/route/user.router"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(router fiber.Router) {
	api := router.Group("/api")
	userRouter.SetupUserRouter(api)
}
