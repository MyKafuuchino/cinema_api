package dto

type CreateCinemaRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Location string `json:"location" validate:"required,max=255"`
}

type UpdateCinemaRequest struct {
	Name     *string `json:"name" validate:"required,max=100"`
	Location *string `json:"location" validate:"required,max=255"`
}
