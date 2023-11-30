package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/http/controllers"
	"github.com/uzzalhcse/amadeus-go/app/http/middleware"
)

func SetUpAuthRoutes(api fiber.Router) {
	AuthController := controllers.NewAuthController()
	api.Post("/login", AuthController.Login)
	api.Post("/register", AuthController.Register)
	api.Post("/forget-password", AuthController.ForgetPassword)

	auth := api.Group("", middleware.Auth())
	auth.Get("/update-profile", AuthController.UpdateProfile)
	auth.Get("/me", AuthController.Me)

}
