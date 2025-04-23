package models

import (
	"time"
)

type User struct {
	ID               int       `gorm:"primaryKey;autoIncrement"`
	Username         string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null"`
	RefreshToken     string
	RefreshTokenExp  time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}