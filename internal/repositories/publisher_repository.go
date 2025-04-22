package repositories

import (
	"bookstore_go_v1/internal/models"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	GetPublisherByID(id int) (*models.Publisher, error)
}

// GetPublisherByID retrieves a publisher by its ID from the database.
func (r *publisherRepository) GetPublisherByID(id int) (*models.Publisher, error) {
	var publisher models.Publisher
	if err := r.db.First(&publisher, id).Error; err != nil {
		return nil, err
	}
	return &publisher, nil
}

type publisherRepository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) PublisherRepository {
	return &publisherRepository{db: db}
}
