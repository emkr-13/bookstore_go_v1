package handlers

import (
	"bookstore_go_v1/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.authService.Register(input.Username, input.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    authToken, refreshToken, err := h.authService.Login(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":        "Login successful",
        "access_token":  authToken,
        "refresh_token": refreshToken,
    })
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
    var input struct {
        RefreshToken string `json:"refresh_token"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    authToken, refreshToken, err := h.authService.RefreshToken(input.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":        "Token refreshed successfully",
        "access_token":  authToken,
        "refresh_token": refreshToken,
    })
}

func (h *AuthHandler) Logout(c *gin.Context) {
    userID := c.MustGet("user_id").(string)
    if err := h.authService.Logout(userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, gin.H{"message": "Logged out successfully"})
}