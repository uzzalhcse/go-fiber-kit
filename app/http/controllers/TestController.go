package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/services"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go"
)

// TestController defines a controller for handling test-related requests
type TestController struct {
	*BaseController
	TestService *services.TestService
	amadeus     *amadeus.Amadeus
}

// NewTestController creates a new instance of the test controller
func NewTestController(testService *services.TestService) *TestController {
	that := NewBaseController()
	return &TestController{
		BaseController: that,
		TestService:    testService,
		amadeus:        amadeus.NewAmadeus(that.Config.Amadeus.APIKey, that.Config.Amadeus.APISecret),
	}
}

func (that *TestController) Test(c *fiber.Ctx) error {
	response, err := that.amadeus.FlightService.OfferSearchRequest.
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
