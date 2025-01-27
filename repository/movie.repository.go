package repository

import (
	"cinema_api/model"
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Create(movie *model.Movie) error
	FindById(id uint) (*model.Movie, error)
	FindAll() ([]model.Movie, error)
	Update(movie *model.Movie) error
	Delete(id uint) error
}

type movieRepository struct {
	db *gorm.DB
}

func (r *movieRepository) Create(movie *model.Movie) error {
	slugResult := slug.Make(movie.Title)
	movie.Slug = slugResult

	if err := r.db.Create(movie).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("movie:%s already exists", movie.Title)
		}
		return err
	}
	return nil
}

func (r *movieRepository) FindById(id uint) (*model.Movie, error) {
	var movie *model.Movie
	if err := r.db.Where("id = ?", id).First(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *movieRepository) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) Update(movie *model.Movie) error {
	slugResult := slug.Make(movie.Title)
	movie.Slug = slugResult
	if err := r.db.Save(movie).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("movie:%s already exists", movie.Title)
		}
		return err
	}
	return nil
}

func (r *movieRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.Movie{}).Error
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}
