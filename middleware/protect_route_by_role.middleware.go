package middleware

import (
	"cinema_api/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

func ProtectRouteByRole(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return fiber.NewError(fiber.StatusForbidden, "No Authorization Token")
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusForbidden, "Invalid Authorization Token")
		}

		claim, err := helper.DecodeJWTToken(tokenParts[1])
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, "Invalid Authorization Token")
		}

		log.Infof("Claim: %v", claim)

		for _, role := range roles {
			if role != claim.Role {
				return fiber.NewError(fiber.StatusForbidden, "Insufficient role permissions")
			}
		}

		ctx.Locals("claim", claim)

		return ctx.Next()
	}
}
