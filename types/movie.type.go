package types

import "time"

type MovieResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	Duration    int       `json:"duration"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
