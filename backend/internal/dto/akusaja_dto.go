package dto

import "time"

type AkusajaResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAkusajaRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAkusajaRequest struct {
	Name string `json:"name"`
}
