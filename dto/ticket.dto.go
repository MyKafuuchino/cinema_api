package dto

type CreateTicketRequest struct {
	UserID      uint   `json:"user_id" validate:"required"`
	ScreeningID uint   `json:"screening_id" validate:"required"`
	SeatNumber  int    `json:"seat_number" validate:"required,gte=0"`
	Status      string `json:"status,omitempty" validate:"omitempty,oneof=booked paid canceled"`
}

type UpdateTicketRequest struct {
	ID         uint `json:"id"`
	SeatNumber *int `json:"seat_number" validate:"omitempty,gte=0"`
}
