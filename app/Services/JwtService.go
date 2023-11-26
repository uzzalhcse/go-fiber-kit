package Services

import "github.com/uzzalhcse/amadeus-go/app/Models"

type JWTService interface {
	GenerateToken(user *Models.User) (string, error)
}
