package repositories

import (
    "bookstore_go_v1/internal/models"
    "gorm.io/gorm"
)

type BookRepository interface {
    CreateBook(book *models.Book) error
    GetAllBooks() ([]models.Book, error)
    GetBookByID(id int) (*models.Book, error)
    UpdateBook(book *models.Book) error
    DeleteBook(id int) error
}

type bookRepository struct {
    db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
    return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(book *models.Book) error {
    return r.db.Create(book).Error
}

func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
    var books []models.Book
    err := r.db.Preload("Author").Preload("Publisher").Find(&books).Error
    return books, err
}

func (r *bookRepository) GetBookByID(id int) (*models.Book, error) {
    var book models.Book
    err := r.db.Preload("Author").Preload("Publisher").First(&book, id).Error
    return &book, err
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
    return r.db.Save(book).Error
}

func (r *bookRepository) DeleteBook(id int) error {
    return r.db.Delete(&models.Book{}, id).Error
}