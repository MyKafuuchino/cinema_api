package repository

import (
	"cinema_api/model"
	"cinema_api/types"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindById(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll(paramsReq *types.QueryParamRequest) ([]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r userRepository) FindById(id uint) (*model.User, error) {
	var user *model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) FindByEmail(email string) (*model.User, error) {
	var user *model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) FindAll(paramsReq *types.QueryParamRequest) ([]model.User, error) {
	var users []model.User
	if err := r.db.Limit(paramsReq.Limit).Offset(paramsReq.Offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r userRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
