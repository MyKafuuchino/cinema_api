package types

import "time"

type TicketResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	ScreeningID uint      `json:"screening_id"`
	SeatNumber  int       `json:"seat_number"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
