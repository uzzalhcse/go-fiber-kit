// Repositories/auth_repository.go

package repositories

import "github.com/uzzalhcse/amadeus-go/app/models"

type AuthRepository interface {
	FindUserByUsername(username string) (*models.User, error)
	FindUserByID(userID string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(username string, updatedUser *models.User) error
}
