package flight

import (
	"encoding/json"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus/flight/models"
)

type SearchRequest struct {
	Service                 *Service
	originLocationCode      string
	destinationLocationCode string
	departureDate           string
	returnDate              string
	adults                  string
	includedAirlineCodes    string
	max                     string
}

func (r *SearchRequest) OriginLocationCode(originLocationCode string) *SearchRequest {
	r.originLocationCode = originLocationCode
	return r
}

func (r *SearchRequest) DestinationLocationCode(destinationLocationCode string) *SearchRequest {
	r.destinationLocationCode = destinationLocationCode
	return r
}

func (r *SearchRequest) DepartureDate(departureDate string) *SearchRequest {
	r.departureDate = departureDate
	return r
}

func (r *SearchRequest) ReturnDate(returnDate string) *SearchRequest {
	r.returnDate = returnDate
	return r
}

func (r *SearchRequest) Adult(adults string) *SearchRequest {
	r.adults = adults
	return r
}

func (r *SearchRequest) IncludedAirlineCodes(includedAirlineCodes string) *SearchRequest {
	r.includedAirlineCodes = includedAirlineCodes
	return r
}

func (r *SearchRequest) Max(max string) *SearchRequest {
	r.max = max
	return r
}

func (r *SearchRequest) Get() (*models.FlightOfferData, error) {
	err := r.Service.Client.GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := r.Service.Client.Client.R().
		SetQueryParams(map[string]string{
			"originLocationCode":      r.originLocationCode,
			"destinationLocationCode": r.destinationLocationCode,
			"departureDate":           r.departureDate,
			"returnDate":              r.returnDate,
			"adults":                  r.adults,
			"includedAirlineCodes":    r.includedAirlineCodes,
			"max":                     r.max,
		}).
		SetHeader("Authorization", "Bearer "+r.Service.Client.AccessToken).
		Get("https://test.api.amadeus.com/v2/shopping/flight-offers")

	if err != nil {
		return nil, err
	}
	var flightOfferData models.FlightOfferData
	if err := json.Unmarshal(resp.Body(), &flightOfferData); err != nil {
		return nil, err
	}

	return &flightOfferData, nil
}
