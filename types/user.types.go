package types

import "time"

type UserResponse struct {
	ID       uint      `json:"id"`
	FullName string    `json:"full_name" `
	Email    string    `json:"email" `
	Role     string    `json:"role"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
