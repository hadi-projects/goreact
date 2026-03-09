package dto

import "time"

type TesttResponse struct {
	ID        uint      `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTesttRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTesttRequest struct {
	Name string `json:"name"`
}
