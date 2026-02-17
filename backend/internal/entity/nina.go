package entity

import "time"

type Nina struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Names string `gorm:"type:varchar(255);not null" json:"names"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
