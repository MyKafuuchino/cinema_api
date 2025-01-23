package repository

import (
	"cinema_api/model"
	"gorm.io/gorm"
)

type CinemaRepository interface {
	Create(cinema *model.Cinema) error
	FindById(id uint) (*model.Cinema, error)
	FindAll() ([]model.Cinema, error)
	Update(cinema *model.Cinema) error
	Delete(id uint) error
}

type cinemaRepository struct {
	db *gorm.DB
}

func (r *cinemaRepository) Create(cinema *model.Cinema) error {
	return r.db.Create(cinema).Error
}

func (r *cinemaRepository) FindById(id uint) (*model.Cinema, error) {
	var cinema *model.Cinema
	if err := r.db.Where("id = ?", id).First(&cinema).Error; err != nil {
		return nil, err
	}
	return cinema, nil
}

func (r *cinemaRepository) FindAll() ([]model.Cinema, error) {
	var cinemas []model.Cinema
	if err := r.db.Find(&cinemas).Error; err != nil {
		return nil, err
	}
	return cinemas, nil
}

func (r *cinemaRepository) Update(cinema *model.Cinema) error {
	return r.db.Save(cinema).Error
}

func (r *cinemaRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.Cinema{}).Error
}

func NewCinemaRepository(db *gorm.DB) CinemaRepository {
	return &cinemaRepository{db: db}
}
