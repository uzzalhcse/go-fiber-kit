package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Http/Controllers"
	"github.com/uzzalhcse/amadeus-go/app/Repositories"
	"github.com/uzzalhcse/amadeus-go/app/Services"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
)

func SetUpAuthRoutes(api fiber.Router) {
	testRepo := Repositories.NewTestRepository(bootstrap.App().DB)
	testService := Services.NewTestService(testRepo)
	testController := Controllers.NewTestController(testService)
	api.Get("/", testController.Test)
	api.Get("/test", testController.GetAllHandler)

}
