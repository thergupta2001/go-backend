package models

import (
	"time"
)

type Patient struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:255;not null"`
	CreatedAt time.Time
	Status string
}