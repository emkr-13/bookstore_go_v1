package services

import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/repositories"
	"errors"
)

type BookService interface {
    CreateBook(book *models.Book) (*models.Book, error)
    GetAllBooks() ([]models.Book, error)
    GetBookByID(id int) (*models.Book, error)
    UpdateBook(book *models.Book) error
    DeleteBook(id int) error
}

type bookService struct {
    bookRepo     repositories.BookRepository
    authorRepo   repositories.AuthorRepository
    publisherRepo repositories.PublisherRepository
}

func NewBookService(
    bookRepo repositories.BookRepository,
    authorRepo repositories.AuthorRepository,
    publisherRepo repositories.PublisherRepository,
) BookService {
    return &bookService{
        bookRepo:     bookRepo,
        authorRepo:   authorRepo,
        publisherRepo: publisherRepo,
    }
}

func (s *bookService) CreateBook(book *models.Book) (*models.Book, error) {
    // Check author exists
    if _, err := s.authorRepo.GetAuthorByID(book.AuthorID); err != nil {
        return nil, errors.New("author not found")
    }

    // Check publisher exists
    if _, err := s.publisherRepo.GetPublisherByID(book.PublisherID); err != nil {
        return nil, errors.New("publisher not found")
    }

    if err := s.bookRepo.CreateBook(book); err != nil {
        return nil, err
    }
    return book, nil
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
    return s.bookRepo.GetAllBooks()
}

func (s *bookService) GetBookByID(id int) (*models.Book, error) {
    return s.bookRepo.GetBookByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
    existingBook, err := s.bookRepo.GetBookByID(book.ID)
    if err != nil {
        return errors.New("book not found")
    }

    existingBook.Title = book.Title
    existingBook.AuthorID = book.AuthorID
    existingBook.PublisherID = book.PublisherID
    existingBook.ISBN = book.ISBN
    existingBook.Price = book.Price
    existingBook.Stock = book.Stock
    existingBook.Description = book.Description
    existingBook.Year = book.Year
    existingBook.Genre = book.Genre

    return s.bookRepo.UpdateBook(existingBook)
}

func (s *bookService) DeleteBook(id int) error {
    if _, err := s.bookRepo.GetBookByID(id); err != nil {
        return errors.New("book not found")
    }

    return s.bookRepo.DeleteBook(id)
}