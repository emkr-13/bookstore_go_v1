package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
    ID        int            `gorm:"primaryKey;autoIncrement"`
    Name      string         `gorm:"not null"`
    Bio       string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}