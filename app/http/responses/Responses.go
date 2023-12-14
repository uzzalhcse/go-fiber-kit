package responses

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Success sends a successful JSON response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"success":     true,
		"status_code": http.StatusOK,
		"data":        data,
		"message":     message,
	})
}

// Error sends a JSON error response
func Error(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success":     false,
		"status_code": http.StatusBadRequest,
		"message":     message,
		//"errors":  nil,
	})
}
