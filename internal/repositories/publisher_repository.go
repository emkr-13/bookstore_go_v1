package repositories

import (
	"bookstore_go_v1/internal/models"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	GetPublisherByID(id int) (*models.Publisher, error)
	CreatePublisher(publisher *models.Publisher) error
	GetAllPublishers() ([]models.Publisher, error)
	UpdatePublisher(publisher *models.Publisher) error
	DeletePublisher(id int) error
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
// CreatePublisher creates a new publisher in the database.
func (r *publisherRepository) CreatePublisher(publisher *models.Publisher) error {
	if err := r.db.Create(publisher).Error; err != nil {
		return err
	}
	return nil
}
// GetAllPublishers retrieves all publishers from the database.
func (r *publisherRepository) GetAllPublishers() ([]models.Publisher, error) {
	var publishers []models.Publisher
	if err := r.db.Find(&publishers).Error; err != nil {
		return nil, err
	}
	return publishers, nil
}
// UpdatePublisher updates an existing publisher in the database.
func (r *publisherRepository) UpdatePublisher(publisher *models.Publisher) error {
	if err := r.db.Save(publisher).Error; err != nil {
		return err
	}
	return nil
}
// DeletePublisher deletes a publisher from the database.
func (r *publisherRepository) DeletePublisher(id int) error {
	if err := r.db.Delete(&models.Publisher{}, id).Error; err != nil {
		return err
	}
	return nil
}
