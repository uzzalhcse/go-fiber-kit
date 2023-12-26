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

	//response, err := that.amadeus.FlightService.OfferSearchRequest.
	//	OriginLocationCode(request.OriginLocationCode).
	//	DestinationLocationCode(request.DestinationLocationCode).
	//	DepartureDate(request.DepartureDate).
	//	ReturnDate(request.ReturnDate).
	//	Adult(request.Adult).
	//	Max(request.Max).
	//	IncludedAirlineCodes(request.IncludedAirlineCodes).
	//	Get()
	response, err := that.amadeus.FlightService.OfferSearchRequest.
		Send(request)

	if err != nil {
		return responses.Error(c, err.Error())
	}
	return responses.Success(c, "Offer Search items", fiber.Map{
		"items": response.Data,
	})
}

// OfferPrice handles the login route
func (that *FlightBookingController) OfferPrice(c *fiber.Ctx) error {
	var request flightrequests.PriceSearchRequest

	// Validate the request
	if err := request.ParseAndValidate(c, &request); err != nil {
		return responses.Error(c, err.Error())
	}

	response, err := that.amadeus.FlightService.PricingRequest.
		FlightOffers(request.FlightOfferData).
		Send()

	if err != nil {
		return responses.Error(c, err.Error())
	}
	return responses.Success(c, "Offer Search items", fiber.Map{
		"items": response,
	})
}

func (that *FlightBookingController) CreateOrder(c *fiber.Ctx) error {
	var request flightrequests.CreateOrderRequest
	response, err := that.amadeus.FlightService.CreateOrder.FlightOffers(request.FlightOfferData).
		Travelers(request.Travelers).
		Remarks(request.Remarks).
		TicketingAgreement(request.TicketingAgreement).
		Contacts(request.Contacts).
		Send()
	if err != nil {
		return err
	}
	return responses.Success(c, "Order Created Successfully", response)
}

func (that *FlightBookingController) Airports(c *fiber.Ctx) error {
	response, err := that.amadeus.FlightService.SearchAirports.
		CountryCode(c.Query("countryCode")).
		SubType(c.Query("subType")).
		Keyword(c.Query("keyword")).
		Get()

	if err != nil {
		return responses.Error(c, err.Error())
	}
	return responses.Success(c, "Airport List", fiber.Map{
		"items": response,
	})
}
