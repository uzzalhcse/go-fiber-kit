// Controllers/auth_controller.go

package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go"
)

type FlightBookingController struct {
	*BaseController
	amadeus *amadeus.Amadeus
}

func NewFlightBookingController() *FlightBookingController {
	that := NewBaseController()

	return &FlightBookingController{
		BaseController: that,
		amadeus:        amadeus.NewAmadeus(that.Config.Amadeus.APIKey, that.Config.Amadeus.APISecret),
	}
}

// OfferSearch handles the login route
func (that *FlightBookingController) OfferSearch(c *fiber.Ctx) error {
	response, err := that.amadeus.FlightService.OfferSearchRequest.
		OriginLocationCode("DEL").
		DestinationLocationCode("LON").
		DepartureDate("2023-12-01").
		ReturnDate("2023-12-15").
		Adult("2").
		//Max("5").
		//IncludedAirlineCodes("TG").
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
