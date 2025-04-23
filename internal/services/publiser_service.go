package services
import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/repositories"
	"errors"
)
type PublisherService interface {
	CreatePublisher(publisher *models.Publisher) (*models.Publisher, error)
	GetAllPublishers() ([]models.Publisher, error)
	GetPublisherByID(id int) (*models.Publisher, error)
	UpdatePublisher(publisher *models.Publisher) error
	DeletePublisher(id int) error
}
type publisherService struct {
	publisherRepo repositories.PublisherRepository
}
func NewPublisherService(publisherRepo repositories.PublisherRepository) PublisherService {
	return &publisherService{
		publisherRepo: publisherRepo,
	}
}
func (s *publisherService) CreatePublisher(publisher *models.Publisher) (*models.Publisher, error) {
	if err := s.publisherRepo.CreatePublisher(publisher); err != nil {
		return nil, err
	}
	return publisher, nil
}
func (s *publisherService) GetAllPublishers() ([]models.Publisher, error) {
	return s.publisherRepo.GetAllPublishers()
}
func (s *publisherService) GetPublisherByID(id int) (*models.Publisher, error) {
	publisher, err := s.publisherRepo.GetPublisherByID(id)
	if err != nil {
		return nil, errors.New("publisher not found")
	}
	return publisher, nil
}
func (s *publisherService) UpdatePublisher(publisher *models.Publisher) error {
	if err := s.publisherRepo.UpdatePublisher(publisher); err != nil {
		return err
	}
	return nil
}
func (s *publisherService) DeletePublisher(id int) error {
	if err := s.publisherRepo.DeletePublisher(id); err != nil {
		return err
	}
	return nil
}