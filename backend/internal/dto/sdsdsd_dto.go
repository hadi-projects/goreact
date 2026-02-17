package dto

import "time"

type SdsdsdResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSdsdsdRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateSdsdsdRequest struct {
	Name string `json:"name"`
}
