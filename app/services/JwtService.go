package services

import "github.com/uzzalhcse/go-fiber-kit/app/models"

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
}
