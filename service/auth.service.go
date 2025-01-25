package service

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/model"
	"cinema_api/repository"
	"cinema_api/types"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(loginReq *dto.LoginRequest) (*types.LoginResponse, error)
	Register(reqUser *dto.CreateUserRequest) (*types.UserResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func (s *authService) Login(loginReq *dto.LoginRequest) (*types.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "email or password is incorrect")
		}
		return nil, err
	}

	if !helper.VerifyPassword(user.Password, loginReq.Password) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "email or password is incorrect")
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

func (s *authService) Register(reqUser *dto.CreateUserRequest) (*types.UserResponse, error) {
	hashedPassword, err := helper.HashPassword(reqUser.Password)
	if err != nil {
		log.Fatalf("Failed to hash password for user %s: %v", reqUser.Email, err)
	}

	reqUser.Password = hashedPassword

	user := model.User{
		FullName: reqUser.FullName,
		Email:    reqUser.Email,
		Password: reqUser.Password,
		Role:     reqUser.Role,
	}

	log.Info(user)

	if err := s.userRepo.Create(&user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusConflict, "email has been used :"+err.Error())
		}
		return nil, err
	}

	createResponse := types.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}

	return &createResponse, nil
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}
