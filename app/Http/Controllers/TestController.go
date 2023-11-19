package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/core/container"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus/flight"
)

// TestController defines a controller for handling test-related requests
type TestController struct {
	*BaseController
}

// NewTestController creates a new instance of the test controller
func NewTestController(container *container.Container) *TestController {
	return &TestController{NewBaseController(container)}
}

func (h *TestController) Test(c *fiber.Ctx) error {
	// Access configuration using the embedded Config instance
	//apiKey := h.Container.App.Config.Amadeus.APIKey
	//apiSecret := h.Container.App.Config.Amadeus.APISecret
	//
	//// Access the TestService
	//result, err := h.Container.TestService.DoSomething()
	//if err != nil {
	//	// handle error
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}
	apiKey := "wGhv7zl6MfpoxRGA9YEA3w8dgcKqVqov"
	apiSecret := "4nccwD5I2O5eh9wy"
	client := amadeus.NewClient(apiKey, apiSecret)
	flightService := flight.NewFlightService(client)
	response, err := flightService.OfferSearch().
		OriginLocationCode("DEL").
		DestinationLocationCode("LON").
		DepartureDate("2023-12-01").
		ReturnDate("2023-12-05").
		Adult("2").
		Max("5").
		IncludedAirlineCodes("TG").
		Get()

	if err != nil {
		fmt.Println("Error:", err)
	}
	return c.JSON(fiber.Map{
		"message": "Hello World",
		"status":  "Success",
		"data":    response,
	})
}

// GetAllHandler handles the route to get all records
func (h *TestController) GetAllHandler(c *fiber.Ctx) error {
	models, err := h.Container.TestService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(models)
}
