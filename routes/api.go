package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/go-fiber-kit/app/http/controllers"
	"github.com/uzzalhcse/go-fiber-kit/app/repositories"
	"github.com/uzzalhcse/go-fiber-kit/app/services"
	"github.com/uzzalhcse/go-fiber-kit/bootstrap"
)

func SetUpApiRoutes(api fiber.Router) {
	testRepo := repositories.NewTestRepository(bootstrap.App().DB)
	testService := services.NewTestService(testRepo)
	testController := controllers.NewTestController(testService)
	api.Get("/", testController.Test)
	api.Get("/test", testController.GetAllHandler)
}
