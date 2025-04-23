package router

import (
	"bookstore_go_v1/internal/handlers"
	"bookstore_go_v1/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	// Public routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh-token", authHandler.RefreshToken)
}

func SetupProtectedRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler, bookHandler *handlers.BookHandler,publishersHandler *handlers.PublisherHandler, authorsHandler *handlers.AuthorHandler) {
	// Protected routes (require authentication)
	r.Use(middlewares.AuthMiddleware())
	r.POST("/logout", authHandler.Logout)
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetAllBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
	r.POST("/publishers", publishersHandler.CreatePublisher)
	r.GET("/publishers", publishersHandler.GetAllPublishers)
	r.GET("/publishers/:id", publishersHandler.GetPublisherByID)
	r.PUT("/publishers/:id", publishersHandler.UpdatePublisher)
	r.DELETE("/publishers/:id", publishersHandler.DeletePublisher)
	r.POST("/authors", authorsHandler.CreateAuthor)
	r.GET("/authors", authorsHandler.GetAllAuthors)
	r.GET("/authors/:id", authorsHandler.GetAuthorByID)
	r.PUT("/authors/:id", authorsHandler.UpdateAuthor)
	r.DELETE("/authors/:id", authorsHandler.DeleteAuthor)
}
