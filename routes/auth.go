package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Http/Controllers"
	middleware "github.com/uzzalhcse/amadeus-go/app/Http/Middleware"
)

func SetUpAuthRoutes(api fiber.Router) {
	AuthController := Controllers.NewAuthController()
	api.Post("/login", AuthController.Login)
	api.Post("/register", AuthController.Register)
	api.Post("/forget-password", AuthController.ForgetPassword)

	auth := api.Group("", middleware.Auth())
	auth.Get("/update-profile", AuthController.UpdateProfile)
	auth.Get("/me", AuthController.Me)

}
