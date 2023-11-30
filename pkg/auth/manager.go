package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/models"
)

type Auth struct {
	user *models.User
	ctx  *fiber.Ctx
}

func NewAuth(c *fiber.Ctx, user *models.User) *Auth {
	return &Auth{ctx: c, user: user}
}

func (that *Auth) User() *models.User {
	return that.user
}
