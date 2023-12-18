package flightrequests

import (
	"github.com/uzzalhcse/amadeus-go/pkg/validator"
)

type PriceSearchRequest struct {
	FlightOfferData []interface{} `json:"flightOffer_data" validate:"required"`
	*validator.Request
}
