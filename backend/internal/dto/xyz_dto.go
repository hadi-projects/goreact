package dto

import "time"

type XyzResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateXyzRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateXyzRequest struct {
	Name string `json:"name"`
}
