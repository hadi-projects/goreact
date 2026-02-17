package dto

import "time"

type MakanResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMakanRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateMakanRequest struct {
	Name string `json:"name"`
}
