package dto

import "time"

type NinaResponse struct {
	ID        uint      `json:"id"`
	Names string `json:"names"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateNinaRequest struct {
	Names string `json:"names" binding:"required"`
}

type UpdateNinaRequest struct {
	Names string `json:"names"`
}
