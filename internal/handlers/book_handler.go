package handlers

import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
    bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
    return &BookHandler{bookService: bookService}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdBook, err := h.bookService.CreateBook(&book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "data":    createdBook,
        "message": "Book created successfully",
    })
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
    books, err := h.bookService.GetAllBooks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":    books,
        "message": "Books retrieved successfully",
    })
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
    id := c.Param("id")
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    book, err := h.bookService.GetBookByID(bookID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":    book,
        "message": "Book retrieved successfully",
    })
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
    id := c.Param("id")
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book.ID = bookID
    if err := h.bookService.UpdateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, gin.H{
        "message": "Book updated successfully",
    })
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    if err := h.bookService.DeleteBook(bookID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, gin.H{
        "message": "Book deleted successfully",
    })
}