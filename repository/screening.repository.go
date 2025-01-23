package repository

import (
	"cinema_api/model"
	"fmt"
	"gorm.io/gorm"
)

type ScreeningRepository interface {
	GetAllScreenings() ([]model.Screening, error)
	GetScreeningByID(id uint) (*model.Screening, error)
	GetScreeningsByCinema(id uint) ([]model.Screening, error)
	GetScreeningsByMovie(id uint) ([]model.Screening, error)
	CreateScreening(screeningReq *model.Screening) error
	UpdateScreening(screeningReq *model.Screening) error
	DeleteScreening(id uint) error
}

type screeningRepository struct {
	db *gorm.DB
}

func (s screeningRepository) GetAllScreenings() ([]model.Screening, error) {
	var screenings []model.Screening
	if err := s.db.Find(&screenings).Error; err != nil {
		return nil, err
	}
	return screenings, nil
}

func (s screeningRepository) GetScreeningByID(id uint) (*model.Screening, error) {
	var screening model.Screening
	if err := s.db.Where("id = ?", id).First(&screening).Error; err != nil {
		return nil, err
	}
	return &screening, nil
}

func (s screeningRepository) GetScreeningsByCinema(id uint) ([]model.Screening, error) {
	var screenings []model.Screening
	if err := s.db.Where("cinema_id = ?", id).Find(&screenings).Error; err != nil {
		return nil, err
	}
	return screenings, nil
}

func (s screeningRepository) GetScreeningsByMovie(id uint) ([]model.Screening, error) {
	var screenings []model.Screening
	if err := s.db.Where("movie_id = ?", id).Find(&screenings).Error; err != nil {
		return nil, err
	}
	return screenings, nil
}

func (s screeningRepository) CreateScreening(screeningReq *model.Screening) error {
	var conflictCount int64
	err := s.db.Model(&model.Screening{}).Where("id = ?", screeningReq.ID).
		Where("cinema_id = ? AND screening_time = ?", screeningReq.CinemaID, screeningReq.ScreeningTime).
		Count(&conflictCount).Error
	if err != nil {
		return err
	}
	if conflictCount > 0 {
		return fmt.Errorf("screening time conflict in the same cinema")
	}
	err = s.db.Create(screeningReq).Error
	if err != nil {
		return err
	}
	return nil
}

func (s screeningRepository) UpdateScreening(screeningReq *model.Screening) error {
	var conflictCount int64
	err := s.db.Model(&model.Screening{}).
		Where("cinema_id = ? AND screening_time = ? AND id != ?", screeningReq.CinemaID, screeningReq.ScreeningTime, screeningReq.ID).
		Count(&conflictCount).Error
	if err != nil {
		return err
	}
	if conflictCount > 0 {
		return fmt.Errorf("screening time conflict in the same cinema")
	}
	err = s.db.Save(screeningReq).Error
	if err != nil {
		return err
	}

	return nil
}

func (s screeningRepository) DeleteScreening(id uint) error {
	err := s.db.Delete(&model.Screening{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func NewScreeningRepository(db *gorm.DB) ScreeningRepository {
	return &screeningRepository{db: db}
}
