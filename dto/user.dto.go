package dto

type CreateUserRequest struct {
	FullName string `json:"full_name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
	Role     string `json:"role,omitempty" validate:"omitempty,oneof=user admin"`
}

type UpdateUserRequest struct {
	Id       uint    `json:"id"`
	FullName *string `json:"full_name,omitempty" validate:"omitempty,min=1,max=100"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=1,max=255"`
	Role     *string `json:"role,omitempty" validate:"omitempty,oneof=user admin"`
}
