package seeder

import (
	"cinema_api/database"
	"cinema_api/helper"
	"cinema_api/model"
	"github.com/gofiber/fiber/v2/log"
)

func UserSeed() {
	var existingAdmin model.User
	err := database.Db.Where("email = ?", "admin@example.com").First(&existingAdmin).Error
	if err == nil {
		log.Info("Admin user already exists with email: %s", existingAdmin.Email)
		return
	}

	admin := model.User{
		Email:    "admin@example.com",
		FullName: "Administrator",
		Password: "admin",
		Role:     "admin",
	}

	hashedPassword, err := helper.HashPassword(admin.Password)
	if err != nil {
		log.Fatalf("Failed to hash password for user %s: %v", admin.Email, err)
	}
	admin.Password = hashedPassword

	err = database.Db.Create(&admin).Error
	if err != nil {
		log.Fatalf("Failed to create user %s: %v", admin.Email, err)
	} else {
		log.Info("User %s created successfully", admin.Email)
	}
}
