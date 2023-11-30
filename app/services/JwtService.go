package services

import "github.com/uzzalhcse/amadeus-go/app/models"

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
}
