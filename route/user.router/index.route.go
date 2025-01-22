package userRouter

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(c fiber.Router) {
	userRouter := c.Group("/users")
	userRouter.Get("", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"success": true,
			"message": "success",
		})
	})
}
