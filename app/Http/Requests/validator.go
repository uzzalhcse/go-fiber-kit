// requests/validator.go

package Requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate validates a request using the go-playground validator
func Validate(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}
