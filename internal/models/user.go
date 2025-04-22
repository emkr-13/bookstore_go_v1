package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username         string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null"`
	RefreshToken     string
	RefreshTokenExp  time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}