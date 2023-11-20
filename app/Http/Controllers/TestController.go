package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Services"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus/flight"
)

// TestController defines a controller for handling test-related requests
type TestController struct {
	*BaseController
	TestService *Services.TestService
}

// NewTestController creates a new instance of the test controller
func NewTestController(testService *Services.TestService) *TestController {
	return &TestController{
		BaseController: NewBaseController(),
		TestService:    testService,
	}
}

func (that *TestController) Test(c *fiber.Ctx) error {
	// Access configuration using the embedded Config instance
	apiKey := that.Config.Amadeus.APIKey
	apiSecret := that.Config.Amadeus.APISecret
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
func (that *TestController) GetAllHandler(c *fiber.Ctx) error {
	models, err := that.TestService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(models)
}
