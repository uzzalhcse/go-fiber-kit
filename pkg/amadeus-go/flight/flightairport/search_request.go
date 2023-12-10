// amadeus-go/flight/airport/search_request.go
package flightairport

import (
	"encoding/json"
	"fmt"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/response"
)

type SearchRequest struct {
	Service     *client.Client
	subType     string
	keyword     string
	countryCode string
}

func NewSearchRequest(client *client.Client) *SearchRequest {
	return &SearchRequest{Service: client}
}

func (r *SearchRequest) SubType(subType string) *SearchRequest {
	r.subType = subType
	return r
}

func (r *SearchRequest) Keyword(keyword string) *SearchRequest {
	r.keyword = keyword
	return r
}

func (r *SearchRequest) CountryCode(countryCode string) *SearchRequest {
	r.countryCode = countryCode
	return r
}

func (r *SearchRequest) Get() (interface{}, error) {
	err := r.Service.GetAccessToken()
	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{
		"subType":     r.subType,
		"keyword":     r.keyword,
		"countryCode": r.countryCode,
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
		Get(r.Service.BaseUrl + "/v1/reference-data/locations")

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

	var result interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
