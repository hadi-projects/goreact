package dto

import "time"

type MinaResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMinaRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateMinaRequest struct {
	Name string `json:"name"`
}
