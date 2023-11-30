// Repositories/auth_repository_impl.go

package repositories

import (
	"errors"
	"github.com/uzzalhcse/amadeus-go/app/models"
	"gorm.io/gorm"
	"log"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}

func (r *AuthRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) CreateUser(user *models.User) error {
	if r.DB == nil {
		return errors.New("DB is nil")
	}

	result := r.DB.Create(user)
	if result.Error != nil {
		log.Println("Error creating user:", result.Error)
		return result.Error
	}

	return nil
}

func (r *AuthRepositoryImpl) UpdateUser(username string, updatedUser *models.User) error {
	if r.DB == nil {
		return errors.New("DB is nil")
	}

	// Assuming you have a primary key field named "id" in your User model
	result := r.DB.Model(&models.User{}).Where("username = ?", username).Updates(updatedUser)
	if result.Error != nil {
		log.Println("Error updating user:", result.Error)
		return result.Error
	}

	return nil
}
func (r *AuthRepositoryImpl) FindUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
