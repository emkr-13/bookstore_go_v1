package repositories

import (
	"bookstore_go_v1/internal/models"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAuthorByID(id int) (*models.Author, error)
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
