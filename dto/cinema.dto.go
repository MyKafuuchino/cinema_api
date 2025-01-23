package dto

type CreateCinemaRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Location string `json:"location" validate:"required,max=255"`
}

type UpdateCinemaRequest struct {
	ID       uint    `json:"id"`
	Name     *string `json:"name,omitempty" validate:"omitempty,max=100"`
	Location *string `json:"location,omitempty" validate:"omitempty,max=255"`
}
