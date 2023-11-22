package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Http/Controllers"
)

func SetUpAuthRoutes(api fiber.Router) {
	AuthController := Controllers.NewAuthController()
	api.Post("/login", AuthController.Login)
	api.Post("/register", AuthController.Register)
	api.Post("/update-profile", AuthController.UpdateProfile)
	api.Post("/forget-password", AuthController.ForgetPassword)

}
