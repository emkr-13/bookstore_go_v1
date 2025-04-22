package repositories

import (
	"bookstore_go_v1/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByRefreshToken(refreshToken string) (*models.User, error)
	UpdateRefreshToken(userID string, refreshToken string, exp int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByRefreshToken(refreshToken string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateRefreshToken(userID string, refreshToken string, exp int64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"refresh_token":     refreshToken,
		"refresh_token_exp": time.Unix(exp, 0),
	}).Error
}
