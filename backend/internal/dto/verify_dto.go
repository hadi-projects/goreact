package dto

import "time"

type VerifyResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"Name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateVerifyRequest struct {
	Name string `json:"Name" binding:"required"`
}

type UpdateVerifyRequest struct {
	Name string `json:"Name"`
}
