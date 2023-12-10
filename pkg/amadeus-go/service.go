package amadeus

import (
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/flight/booking"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/flight/flightairport"
)

type FlightService struct {
	OfferSearchRequest *booking.OfferSearchRequest
	PricingRequest     *booking.PricingRequest
	SearchAirports     *flightairport.SearchRequest
}

type HotelService struct {
	//
}

type TransportService struct {
	//
}
