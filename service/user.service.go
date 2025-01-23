package service

import (
	"cinema_api/dto"
	"cinema_api/helper"
	"cinema_api/model"
	"cinema_api/repository"
	"cinema_api/types"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers() ([]types.UserResponse, error)
	GetUserById(id uint) (*types.UserResponse, error)
	CreateUser(createUserReq *dto.CreateUserRequest) (*types.UserResponse, error)
	UpdateUserById(id uint, updateRequest *dto.UpdateUserRequest) (*types.UserResponse, error)
	DeleteUserById(id uint) (*types.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s userService) GetAllUsers() ([]types.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	allUsersResponse := make([]types.UserResponse, len(users))
	for i, user := range users {
		allUsersResponse[i] = types.UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
			Role:     user.Role,
			CreateAt: user.CreatedAt,
			UpdateAt: user.UpdatedAt,
		}
	}
	return allUsersResponse, nil
}

func (s userService) GetUserById(id uint) (*types.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return nil, err
	}

	userResponse := types.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}

	return &userResponse, nil
}

func (s userService) CreateUser(createUserReq *dto.CreateUserRequest) (*types.UserResponse, error) {
	hashedPassword, err := helper.HashPassword(createUserReq.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		FullName: createUserReq.FullName,
		Email:    createUserReq.Email,
		Password: hashedPassword,
		Role:     createUserReq.Role,
	}

	if err := s.userRepo.Create(user); err != nil {
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "User already exists : "+err.Error())
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal server error : "+err.Error())
	}

	userResponse := types.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}

	return &userResponse, nil
}

func (s userService) UpdateUserById(id uint, updateRequest *dto.UpdateUserRequest) (*types.UserResponse, error) {
	user, err := s.userRepo.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return nil, err
	}

	helper.UpdateFields(user, updateRequest)

	if updateRequest.Password != nil {
		hashedPassword, err := helper.HashPassword(*updateRequest.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update user : "+err.Error())
	}

	userResponse := types.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}

	return &userResponse, nil
}

func (s userService) DeleteUserById(id uint) (*types.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return nil, err
	}

	if err := s.userRepo.Delete(user.ID); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete user : "+err.Error())
	}

	userResponse := types.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}
	return &userResponse, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
