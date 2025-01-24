package dto

import "time"

type CreateScreeningRequest struct {
	MovieID       uint      `json:"movie_id" validate:"required"`
	CinemaID      uint      `json:"cinema_id" validate:"required"`
	ScreeningTime time.Time `json:"screening_time" validate:"required"`
	Price         float64   `json:"price" validate:"required"`
}

type UpdateScreeningRequest struct {
	ID            uint       `json:"id"`
	MovieID       *uint      `json:"movie_id,omitempty" validate:"omitempty"`
	CinemaID      *uint      `json:"cinema_id,omitempty" validate:"omitempty"`
	ScreeningTime *time.Time `json:"screening_time,omitempty" validate:"omitempty"`
	Price         *float64   `json:"price,omitempty" validate:"omitempty"`
}
