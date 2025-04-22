package models

import (
	"time"

	"gorm.io/gorm"
)

type Publisher struct {
    ID          int            `gorm:"primaryKey;autoIncrement"`
    Name        string         `gorm:"not null"`
    Address     string         `gorm:"size:500"`
    Description string         `gorm:"size:500"`
    Phone       string         `gorm:"size:20"`
    Email       string         `gorm:"size:255"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}