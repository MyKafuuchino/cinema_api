package dto

type CreateTicketRequest struct {
	UserID      uint   `json:"user_id" validate:"required"`
	ScreeningID uint   `json:"screening_id" validate:"required"`
	SeatNumber  string `json:"seat_number" validate:"required"`
	Status      string `json:"status,omitempty" validate:"omitempty,oneof=booked paid"`
}

type UpdateTicketRequest struct {
	ID          uint    `json:"id"`
	UserID      *uint   `json:"user_id" validate:"omitempty"`
	ScreeningID *uint   `json:"screening_id" validate:"omitempty"`
	SeatNumber  *string `json:"seat_number" validate:"omitempty"`
	Status      *string `json:"status,omitempty" validate:"omitempty,oneof=booked paid"`
}
