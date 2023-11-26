// Repositories/auth_repository_impl.go

package Repositories

import (
	"errors"
	"github.com/uzzalhcse/amadeus-go/app/Models"
	"gorm.io/gorm"
	"log"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}

func (r *AuthRepositoryImpl) FindUserByUsername(username string) (*Models.User, error) {
	var user Models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) CreateUser(user *Models.User) error {
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

func (r *AuthRepositoryImpl) UpdateUser(username string, updatedUser *Models.User) error {
	if r.DB == nil {
		return errors.New("DB is nil")
	}

	// Assuming you have a primary key field named "id" in your User model
	result := r.DB.Model(&Models.User{}).Where("username = ?", username).Updates(updatedUser)
	if result.Error != nil {
		log.Println("Error updating user:", result.Error)
		return result.Error
	}

	return nil
}
func (r *AuthRepositoryImpl) FindUserByID(userID string) (*Models.User, error) {
	var user Models.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
