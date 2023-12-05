// Controllers/auth_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	flightrequests "github.com/uzzalhcse/amadeus-go/app/http/requests/flight"
	"github.com/uzzalhcse/amadeus-go/app/http/responses"
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
	var request flightrequests.FlightSearchRequest

	// Validate the request
	if err := request.ParseAndValidate(c, &request); err != nil {
		return responses.Error(c, err.Error())
	}

	response, err := that.amadeus.FlightService.OfferSearchRequest.
		OriginLocationCode(request.OriginLocationCode).
		DestinationLocationCode(request.DestinationLocationCode).
		DepartureDate(request.DepartureDate).
		ReturnDate(request.ReturnDate).
		Adult(request.Adult).
		Max(request.Max).
		IncludedAirlineCodes(request.IncludedAirlineCodes).
		Get()

	if err != nil {
		return responses.Error(c, err.Error())
	}
	return responses.Success(c, response)
}
