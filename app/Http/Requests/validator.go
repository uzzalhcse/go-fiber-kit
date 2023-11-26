// requests/validator.go

package Requests

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate validates a request using the go-playground validator
func Validate(c *fiber.Ctx, req interface{}) error {
	// Parse the request body
	if err := c.BodyParser(req); err != nil {
		return err
	}

	// Validate the request using the initialized validator
	if err := validate.Struct(req); err != nil {
		// If validation fails, construct a human-readable error message
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := e.Field()
			tagName := e.Tag()
			errorMessage := e.Param()
			fmt.Println(e)

			// Customize error messages based on validation tags
			switch tagName {
			case "required":
				validationErrors = append(validationErrors, fmt.Sprintf("%s is required", fieldName))
			case "min":
				validationErrors = append(validationErrors, fmt.Sprintf("%s must be at least %s characters", fieldName, errorMessage))
			case "max":
				validationErrors = append(validationErrors, fmt.Sprintf("%s cannot be longer than %s characters", fieldName, errorMessage))
			case "email":
				validationErrors = append(validationErrors, fmt.Sprintf("%s must be a valid email address", fieldName))
			default:
				validationErrors = append(validationErrors, fmt.Sprintf("%s is not valid", fieldName))
			}
		}

		// Join validation error messages
		errorMessage := strings.Join(validationErrors, ", ")

		// Return an error with the formatted validation errors
		return fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	// If validation succeeds, return nil
	return nil
}
