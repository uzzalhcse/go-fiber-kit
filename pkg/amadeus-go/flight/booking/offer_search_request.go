// amadeus-go/flight/booking/offer_serarch_requet.go
package booking

import (
	"encoding/json"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/flight/models"
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

	resp, err := r.Service.Client.R().
		SetQueryParams(map[string]string{
			"originLocationCode":      r.originLocationCode,
			"destinationLocationCode": r.destinationLocationCode,
			"departureDate":           r.departureDate,
			"returnDate":              r.returnDate,
			"adults":                  r.adults,
			"includedAirlineCodes":    r.includedAirlineCodes,
			"max":                     r.max,
		}).
		SetHeader("Authorization", "Bearer "+r.Service.AccessToken).
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