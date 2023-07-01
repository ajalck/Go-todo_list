package model

import (
	"time"
)

type Todo struct {
	ID          uint      `gorm:"primary_key" json:"omitempty"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	DeletedAt   time.Time `gorm:"default:current_timestamp"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due         time.Time `json:"due"`
}
