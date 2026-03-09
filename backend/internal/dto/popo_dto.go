package dto

import "time"

type PopoResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePopoRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePopoRequest struct {
	Name string `json:"name"`
}
