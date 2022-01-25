package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model

	ID          uint      `gorm:"primaryKey" json:"ID"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}