package amadeus

import (
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/flight/booking"
)

type Amadeus struct {
	*client.Client
	FlightService *FlightService
}

func NewAmadeus(apiKey, apiSecret string) *Amadeus {
	c := client.NewClient(apiKey, apiSecret)
	return &Amadeus{
		Client: c,
		FlightService: &FlightService{
			OfferSearchRequest: booking.NewOfferSearchRequest(c),
			PricingRequest:     booking.NewPricingRequest(c),
		},
	}
}
