package helper

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "")
	}
	return string(bytes), nil
}

func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
