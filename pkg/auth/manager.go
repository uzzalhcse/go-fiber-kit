package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Models"
)

type Auth struct {
	user *Models.User
	ctx  *fiber.Ctx
}

func NewAuth(c *fiber.Ctx, user *Models.User) *Auth {
	return &Auth{ctx: c, user: user}
}

func (that *Auth) User() *Models.User {
	return that.user
}
