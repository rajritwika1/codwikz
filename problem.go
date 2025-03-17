package models

import (
	"time"

	"gorm.io/gorm"
)

type Problem struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"unique;not null"`
	Description string         `json:"description" gorm:"not null"`
	Difficulty  string         `json:"difficulty" gorm:"not null"`
	Tags        string         `json:"tags"` // e.g., "Arrays, DP, Graphs"
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
