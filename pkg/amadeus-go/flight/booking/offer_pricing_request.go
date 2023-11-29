// amadeus-go/flight/booking/pricing_request.go
package booking

import (
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/flight/models"
)

type PricingRequest struct {
	Service      *client.Client
	flightOffers models.FlightOffer
}

func NewPricingRequest(client *client.Client) *PricingRequest {
	return &PricingRequest{Service: client}
}

func (r *PricingRequest) FlightOffers(flightOffers models.FlightOffer) *PricingRequest {
	r.flightOffers = flightOffers
	return r
}

func (r *PricingRequest) Send() (string, error) {
	err := r.Service.GetAccessToken()
	if err != nil {
		return "", err
	}

	requestBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type":         "flight-offers-pricing",
			"flightOffers": r.flightOffers,
		},
	}

	resp, err := r.Service.Client.R().
		SetBody(requestBody).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+r.Service.AccessToken).
		Post("https://test.api.amadeus.com/v1/shopping/flight-offers/pricing")

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
