// amadeus-go/flight/booking/offer_serarch_requet.go
package booking

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/models"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/response"
	"net/http"
)

type OfferSearchRequest struct {
	Service                 *client.Client
	originLocationCode      string
	destinationLocationCode string
	departureDate           string
	returnDate              string
	adults                  string
	includedAirlineCodes    string
	max                     string
}

func NewOfferSearchRequest(client *client.Client) *OfferSearchRequest {
	return &OfferSearchRequest{Service: client}
}

func (r *OfferSearchRequest) OriginLocationCode(originLocationCode string) *OfferSearchRequest {
	r.originLocationCode = originLocationCode
	return r
}

func (r *OfferSearchRequest) DestinationLocationCode(destinationLocationCode string) *OfferSearchRequest {
	r.destinationLocationCode = destinationLocationCode
	return r
}

func (r *OfferSearchRequest) DepartureDate(departureDate string) *OfferSearchRequest {
	r.departureDate = departureDate
	return r
}

func (r *OfferSearchRequest) ReturnDate(returnDate string) *OfferSearchRequest {
	r.returnDate = returnDate
	return r
}

func (r *OfferSearchRequest) Adult(adults string) *OfferSearchRequest {
	r.adults = adults
	return r
}

func (r *OfferSearchRequest) IncludedAirlineCodes(includedAirlineCodes string) *OfferSearchRequest {
	r.includedAirlineCodes = includedAirlineCodes
	return r
}

func (r *OfferSearchRequest) Max(max string) *OfferSearchRequest {
	r.max = max
	return r
}

func (r *OfferSearchRequest) Get() (*models.FlightOfferData, error) {
	err := r.Service.GetAccessToken()
	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{
		"originLocationCode":      r.originLocationCode,
		"destinationLocationCode": r.destinationLocationCode,
		"departureDate":           r.departureDate,
		"returnDate":              r.returnDate,
		"adults":                  r.adults,
		"includedAirlineCodes":    r.includedAirlineCodes,
		"max":                     r.max,
	}

	// Remove empty query parameters
	for key, value := range queryParams {
		if value == "" {
			delete(queryParams, key)
		}
	}

	resp, err := r.Service.Client.R().
		SetQueryParams(queryParams).
		SetHeader("Authorization", "Bearer "+r.Service.AccessToken).
		Get("https://test.api.amadeus.com/v2/shopping/flight-offers")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 400 {
		var errorResponse response.ErrorResponse
		if err := json.Unmarshal(resp.Body(), &errorResponse); err != nil {
			return nil, err
		}
		fmt.Println("errorResponse", errorResponse.Errors[0])

		// Assuming the first error detail contains the relevant information
		if len(errorResponse.Errors) > 0 {
			return nil, fmt.Errorf("API error: %s", errorResponse.Errors[0].Detail)
		}
	}

	var flightOfferData models.FlightOfferData
	if err := json.Unmarshal(resp.Body(), &flightOfferData); err != nil {
		return nil, err
	}

	return &flightOfferData, nil
}
func (r *OfferSearchRequest) Send(requestBody interface{}) (*models.FlightOfferData, error) {
	err := r.Service.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	resp, err := r.Service.Client.R().
		SetBody(requestBody).
		SetHeader("Authorization", "Bearer "+r.Service.AccessToken).
		Post("https://test.api.amadeus.com/v2/shopping/flight-offers")

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK {
		var errorResponse response.ErrorResponse
		if err := json.Unmarshal(resp.Body(), &errorResponse); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}

		if len(errorResponse.Errors) > 0 {
			return nil, fmt.Errorf("API error (Status %d): %s", statusCode, errorResponse.Errors[0].Detail)
		}

		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	if len(resp.Body()) == 0 {
		return nil, errors.New("empty response body")
	}

	var flightOfferData models.FlightOfferData
	if err := json.Unmarshal(resp.Body(), &flightOfferData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &flightOfferData, nil
}
func (r *OfferSearchRequest) GetRaw(requestBody interface{}) (interface{}, error) {
	err := r.Service.GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := r.Service.Client.R().
		SetBody(requestBody).
		SetHeader("Authorization", "Bearer "+r.Service.AccessToken).
		Post("https://test.api.amadeus.com/v2/shopping/flight-offers")

	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
