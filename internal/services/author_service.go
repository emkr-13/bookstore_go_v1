package services
import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/repositories"
	"errors"
)
type AuthorService interface {
	CreateAuthor(author *models.Author) (*models.Author, error)
	GetAllAuthors() ([]models.Author, error)
	GetAuthorByID(id int) (*models.Author, error)
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id int) error
}
type authorService struct {
	authorRepo repositories.AuthorRepository
}
func NewAuthorService(authorRepo repositories.AuthorRepository) AuthorService {
	return &authorService{
		authorRepo: authorRepo,
	}
}
func (s *authorService) CreateAuthor(author *models.Author) (*models.Author, error) {
	if err := s.authorRepo.CreateAuthor(author); err != nil {
		return nil, err
	}
	return author, nil
}
func (s *authorService) GetAllAuthors() ([]models.Author, error) {
	return s.authorRepo.GetAllAuthors()
}
func (s *authorService) GetAuthorByID(id int) (*models.Author, error) {
	author, err := s.authorRepo.GetAuthorByID(id)
	if err != nil {
		return nil, errors.New("author not found")
	}
	return author, nil
}
func (s *authorService) UpdateAuthor(author *models.Author) error {
	if err := s.authorRepo.UpdateAuthor(author); err != nil {
		return err
	}
	return nil
}
func (s *authorService) DeleteAuthor(id int) error {
	if err := s.authorRepo.DeleteAuthor(id); err != nil {
		return err
	}
	return nil
}