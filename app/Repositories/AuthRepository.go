// Repositories/auth_repository.go

package Repositories

import "github.com/uzzalhcse/amadeus-go/app/Models"

type AuthRepository interface {
	FindUserByUsername(username string) (*Models.User, error)
	FindUserByID(userID string) (*Models.User, error)
	CreateUser(user *Models.User) error
	UpdateUser(username string, updatedUser *Models.User) error
}
