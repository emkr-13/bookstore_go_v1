package handlers

import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService services.AuthorService
}
func NewAuthorHandler(authorService services.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAuthor, err := h.authorService.CreateAuthor(&author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    createdAuthor,
		"message": "Author created successfully",
	})
}
func (h *AuthorHandler) GetAllAuthors(c *gin.Context) {
	authors, err := h.authorService.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    authors,
		"message": "Authors retrieved successfully",
	})
}
func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	author, err := h.authorService.GetAuthorByID(authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if author == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    author,
		"message": "Author retrieved successfully",
	})
}
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author.ID = authorID
	if err := h.authorService.UpdateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    author,
		"message": "Author updated successfully",
	})
}
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	if err := h.authorService.DeleteAuthor(authorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Author deleted successfully",
	})
}