package helper

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func StringToUint(stringNumber string) (uint, error) {
	convertedUint, err := strconv.ParseUint(stringNumber, 10, 64)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "String to Uint conversion failed :"+err.Error())
	}
	return uint(convertedUint), nil
}
