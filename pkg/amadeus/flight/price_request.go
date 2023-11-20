package flight

import (
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus/flight/models"
)

type PricingRequest struct {
	Service      *Service
	flightOffers models.FlightOffer
}

func (r *PricingRequest) FlightOffers(flightOffers models.FlightOffer) *PricingRequest {
	r.flightOffers = flightOffers
	return r
}

func (r *PricingRequest) Send() (string, error) {
	err := r.Service.Client.GetAccessToken()
	if err != nil {
		return "", err
	}

	requestBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type":         "flight-offers-pricing",
			"flightOffers": r.flightOffers,
		},
	}

	resp, err := r.Service.Client.Client.R().
		SetBody(requestBody).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+r.Service.Client.AccessToken).
		Post("https://test.api.amadeus.com/v1/shopping/flight-offers/pricing")

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
