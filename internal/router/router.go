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

func SetupProtectedRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler, bookHandler *handlers.BookHandler) {
	// Protected routes (require authentication)
	r.Use(middlewares.AuthMiddleware())
	r.POST("/logout", authHandler.Logout)
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetAllBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
}
