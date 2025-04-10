package model

import "time"

type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
