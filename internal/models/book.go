package models

import (
	"time"
)

type Book struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"not null"`
	AuthorID    int       `gorm:"not null"`
	PublisherID int       `gorm:"not null"`
	ISBN        string    `gorm:"unique;not null"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	Description string
	Year        int
	Genre       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}