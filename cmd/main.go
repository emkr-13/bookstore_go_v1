package main

import (
	"bookstore_go_v1/internal/config"
	"bookstore_go_v1/internal/handlers"
	"bookstore_go_v1/internal/middlewares"
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/repositories"
	"bookstore_go_v1/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    // Auto migrate database models
    err = cfg.DB.AutoMigrate(
        &models.User{},
        &models.Author{},
        &models.Publisher{},
        &models.Book{},
    )
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    // Initialize repositories
    userRepo := repositories.NewUserRepository(cfg.DB)
    authorRepo := repositories.NewAuthorRepository(cfg.DB) // Fixed
    publisherRepo := repositories.NewPublisherRepository(cfg.DB) // Fixed
    bookRepo := repositories.NewBookRepository(cfg.DB)

    // Initialize services
    authService := services.NewAuthService(userRepo, cfg.JWTSecret, cfg.AuthExp, cfg.RefreshExp)
    bookService := services.NewBookService(bookRepo, authorRepo, publisherRepo)

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)
    bookHandler := handlers.NewBookHandler(bookService)

    // Setup Gin router
    r := gin.Default()

    // Public routes
    public := r.Group("/api/v1")
    {
        public.POST("/register", authHandler.Register)
        public.POST("/login", authHandler.Login)
        public.POST("/refresh-token", authHandler.RefreshToken)
    }

    // Protected routes (require authentication)
    protected := r.Group("/api/v1")
    protected.Use(middlewares.AuthMiddleware())
    {
        protected.POST("/logout", authHandler.Logout)
        protected.POST("/books", bookHandler.CreateBook)
        protected.GET("/books", bookHandler.GetAllBooks)
        protected.GET("/books/:id", bookHandler.GetBookByID)
        protected.PUT("/books/:id", bookHandler.UpdateBook)
        protected.DELETE("/books/:id", bookHandler.DeleteBook)
    }

    // Start server
    log.Printf("Server running on port %s", cfg.AppPort)
    if err := r.Run(":" + cfg.AppPort); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}