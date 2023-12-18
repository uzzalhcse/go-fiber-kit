// amadeus-go/flight/booking/pricing_request.go

package booking

import (
	"encoding/json"
	"fmt"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/response"
)

type PricingRequest struct {
	Service      *client.Client
	flightOffers []interface{}
}

func NewPricingRequest(client *client.Client) *PricingRequest {
	return &PricingRequest{Service: client}
}

func (r *PricingRequest) FlightOffers(flightOffers []interface{}) *PricingRequest {
	r.flightOffers = flightOffers
	return r
}

func (r *PricingRequest) Send() (interface{}, error) {
	err := r.Service.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %v", err)
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
		SetHeader("X-HTTP-Method-Override", "GET").
		SetHeader("Authorization", "Bearer "+r.Service.AccessToken).
		Post(r.Service.BaseUrl + "/v1/shopping/flight-offers/pricing")

	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	if resp.StatusCode() != 200 {
		var errorResponse response.ErrorResponse
		if err := json.Unmarshal(resp.Body(), &errorResponse); err != nil {
			return nil, fmt.Errorf("failed to parse error response: %v", err)
		}

		if len(errorResponse.Errors) > 0 {
			return nil, fmt.Errorf("API error: %s", errorResponse.Errors[0].Detail)
		}
	}

	var result interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
