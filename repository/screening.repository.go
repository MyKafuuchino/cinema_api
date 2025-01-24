package repository

import (
	"cinema_api/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ScreeningRepository interface {
	FindAll() ([]model.Screening, error)
	FindById(id uint) (*model.Screening, error)
	FindByCinema(id uint) ([]model.Screening, error)
	FindByMovie(id uint) ([]model.Screening, error)
	Create(screeningReq *model.Screening) error
	Update(screeningReq *model.Screening) error
	Delete(id uint) error
}

type screeningRepository struct {
	db *gorm.DB
}

func (r *screeningRepository) FindAll() ([]model.Screening, error) {
	var screenings []model.Screening
	if err := r.db.Find(&screenings).Error; err != nil {
		return nil, err
	}
	return screenings, nil
}

func (r *screeningRepository) FindById(id uint) (*model.Screening, error) {
	var screening model.Screening
	if err := r.db.Where("id = ?", id).First(&screening).Error; err != nil {
		return nil, err
	}
	return &screening, nil
}

func (r *screeningRepository) FindByCinema(id uint) ([]model.Screening, error) {
	var screenings []model.Screening
	if err := r.db.Where("cinema_id = ?", id).Find(&screenings).Error; err != nil {
		return nil, err
	}
	return screenings, nil
}

func (r *screeningRepository) FindByMovie(id uint) ([]model.Screening, error) {
	var screenings []model.Screening
	if err := r.db.Where("movie_id = ?", id).Find(&screenings).Error; err != nil {
		return nil, err
	}
	return screenings, nil
}

func (r *screeningRepository) Create(screeningReq *model.Screening) error {
	var conflictCount int64
	err := r.db.Model(&model.Screening{}).
		Where("cinema_id = ? AND screening_time = ?", screeningReq.CinemaID, screeningReq.ScreeningTime).
		Count(&conflictCount).Error
	if err != nil {
		return err
	}
	if conflictCount > 0 {
		return fmt.Errorf("screening time conflict in the same cinema")
	}
	err = r.db.Create(screeningReq).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *screeningRepository) Update(screeningReq *model.Screening) error {
	existingScreening, err := r.FindById(screeningReq.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("screening not found")
		}
		return err
	}

	if existingScreening.CinemaID == screeningReq.CinemaID && existingScreening.ScreeningTime.Equal(screeningReq.ScreeningTime) {
		return r.db.Save(screeningReq).Error
	}

	var conflictCount int64
	err = r.db.Model(&model.Screening{}).
		Where("cinema_id = ? AND screening_time = ? AND id != ?", screeningReq.CinemaID, screeningReq.ScreeningTime, screeningReq.ID).
		Count(&conflictCount).Error
	if err != nil {
		return err
	}
	if conflictCount > 0 {
		return fmt.Errorf("screening time conflict in the same cinema")
	}

	err = r.db.Save(screeningReq).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *screeningRepository) Delete(id uint) error {
	err := r.db.Delete(&model.Screening{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func NewScreeningRepository(db *gorm.DB) ScreeningRepository {
	return &screeningRepository{db: db}
}
