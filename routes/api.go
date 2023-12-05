package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/http/controllers"
	"github.com/uzzalhcse/amadeus-go/app/repositories"
	"github.com/uzzalhcse/amadeus-go/app/services"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
)

func SetUpApiRoutes(api fiber.Router) {
	testRepo := repositories.NewTestRepository(bootstrap.App().DB)
	testService := services.NewTestService(testRepo)
	testController := controllers.NewTestController(testService)
	api.Get("/", testController.Test)
	api.Get("/test", testController.GetAllHandler)
	flight := api.Group("/flight")

	flightController := controllers.NewFlightBookingController()
	flight.Get("/offer-search", flightController.OfferSearch)
}
