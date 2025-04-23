package handlers

import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type PublisherHandler struct {
	publisherService services.PublisherService
}
func NewPublisherHandler(publisherService services.PublisherService) *PublisherHandler {
	return &PublisherHandler{publisherService: publisherService}
}
func (h *PublisherHandler) CreatePublisher(c *gin.Context) {
	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPublisher, err := h.publisherService.CreatePublisher(&publisher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    createdPublisher,
		"message": "Publisher created successfully",
	})
}
func (h *PublisherHandler) GetAllPublishers(c *gin.Context) {
	publishers, err := h.publisherService.GetAllPublishers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    publishers,
		"message": "Publishers retrieved successfully",
	})
}
func (h *PublisherHandler) GetPublisherByID(c *gin.Context) {
	id := c.Param("id")
	publisherID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publisher ID"})
		return
	}

	publisher, err := h.publisherService.GetPublisherByID(publisherID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if publisher == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    publisher,
		"message": "Publisher retrieved successfully",
	})
}
func (h *PublisherHandler) UpdatePublisher(c *gin.Context) {
	id := c.Param("id")
	publisherID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publisher ID"})
		return
	}

	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publisher.ID = publisherID
	if err := h.publisherService.UpdatePublisher(&publisher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    publisher,
		"message": "Publisher updated successfully",
	})
}
func (h *PublisherHandler) DeletePublisher(c *gin.Context) {
	id := c.Param("id")
	publisherID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publisher ID"})
		return
	}

	if err := h.publisherService.DeletePublisher(publisherID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Publisher deleted successfully",
	})
}