package model

import (
	"time"
)

type Todo struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   *time.Time `gorm:"default:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"default:current_timestamp" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `gorm:"default:current_timestamp" json:"deleted_at,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Due         time.Time  `json:"due"`
}
