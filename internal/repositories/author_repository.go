package repositories

import (
	"bookstore_go_v1/internal/models"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAuthorByID(id int) (*models.Author, error)
	CreateAuthor(author *models.Author) error
	GetAllAuthors() ([]models.Author, error)
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id int) error
	// Add any other methods you need for author management
}

type authorRepository struct {
	db *gorm.DB
}

// GetAuthorByID implements AuthorRepository.
func (a *authorRepository) GetAuthorByID(id int) (*models.Author, error) {
	var author models.Author
	if err := a.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}
// CreateAuthor creates a new author in the database.
func (a *authorRepository) CreateAuthor(author *models.Author) error {
	if err := a.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}
// GetAllAuthors retrieves all authors from the database.
func (a *authorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	if err := a.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
// UpdateAuthor updates an existing author in the database.
func (a *authorRepository) UpdateAuthor(author *models.Author) error {
	if err := a.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}
// DeleteAuthor deletes an author from the database.
func (a *authorRepository) DeleteAuthor(id int) error {
	if err := a.db.Delete(&models.Author{}, id).Error; err != nil {
		return err
	}
	return nil
}
