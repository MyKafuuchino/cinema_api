package service

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/repository"
	"cinema_api/types"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(loginReq *dto.LoginRequest) (*types.LoginResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func (s authService) Login(loginReq *dto.LoginRequest) (*types.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "email or password is incorrect")
		}
		return nil, err
	}

	if err := helper.VerifyPassword(user.Password, loginReq.Password); err {
		return nil, fiber.NewError(fiber.StatusConflict, "email or password is incorrect")
	}

	userPayload := types.UserPayload{
		Id:   user.ID,
		Role: user.Role,
	}

	token, err := helper.GenerateJWTToken(userPayload)

	authResponse := types.LoginResponse{
		Email: user.Email,
		Token: token,
	}

	return &authResponse, err
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}
