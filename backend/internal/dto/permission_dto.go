package dto

import "time"

type PermissionResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePermissionRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePermissionRequest struct {
	Name string `json:"name" binding:"required"`
}
