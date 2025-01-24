package types

import (
	"time"
)

type ScreeningResponse struct {
	ID            uint      `json:"id"`
	MovieID       uint      `json:"movie_id"`
	CinemaID      uint      `json:"cinema_id"`
	ScreeningTime time.Time `json:"screening_time"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
