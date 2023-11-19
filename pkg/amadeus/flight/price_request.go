package flight

type PricingRequest struct {
	service      *Service
	flightOffers FlightOffer
}

func (r *PricingRequest) FlightOffers(flightOffers FlightOffer) *PricingRequest {
	r.flightOffers = flightOffers
	return r
}

func (r *PricingRequest) Send() (string, error) {
	err := r.service.client.GetAccessToken()
	if err != nil {
		return "", err
	}

	requestBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type":         "flight-offers-pricing",
			"flightOffers": r.flightOffers,
		},
	}

	resp, err := r.service.client.Client.R().
		SetBody(requestBody).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+r.service.client.AccessToken).
		Post("https://test.api.amadeus.com/v1/shopping/flight-offers/pricing")

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
