package dto

import "time"

type CreateMovieRequest struct {
	Title       string    `json:"title" validate:"required,max=100"`
	Description string    `json:"description" validate:"required"`
	Genre       string    `json:"genre" validate:"required,max=50"`
	Duration    int       `json:"duration" validate:"required,min=0"`
	ReleaseDate time.Time `json:"release_date" validate:"required"`
}

type UpdateMovieRequest struct {
	ID          uint       `json:"id"`
	Title       *string    `json:"title,omitempty" validate:"omitempty,max=100"`
	Description *string    `json:"description,omitempty" validate:"omitempty"`
	Genre       *string    `json:"genre,omitempty" validate:"omitempty,max=50"`
	Duration    *int       `json:"duration,omitempty" validate:"omitempty,min=0"`
	ReleaseDate *time.Time `json:"release_date,omitempty" validate:"omitempty"`
}
