// amadeus-go/flight/booking/offer_serarch_requet.go
package booking

import (
	"encoding/json"
	"fmt"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/client"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/models"
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/response"
)

type CreateOrderRequest struct {
	service            *client.Client
	flightOffers       []interface{}
	travelers          []models.Traveler
	remarks            models.Remarks
	ticketingAgreement models.TicketingAgreement
	contacts           []models.Contact
}

func NewCreateOrderRequest(client *client.Client) *CreateOrderRequest {
	return &CreateOrderRequest{service: client}
}

func (r *CreateOrderRequest) FlightOffers(flightOffers []interface{}) *CreateOrderRequest {
	r.flightOffers = flightOffers
	return r
}

func (r *CreateOrderRequest) Travelers(travelers []models.Traveler) *CreateOrderRequest {
	r.travelers = travelers
	return r
}

func (r *CreateOrderRequest) Remarks(remarks models.Remarks) *CreateOrderRequest {
	r.remarks = remarks
	return r
}

func (r *CreateOrderRequest) TicketingAgreement(ticketingAgreement models.TicketingAgreement) *CreateOrderRequest {
	r.ticketingAgreement = ticketingAgreement
	return r
}

func (r *CreateOrderRequest) Contacts(contacts []models.Contact) *CreateOrderRequest {
	r.contacts = contacts
	return r
}

func (r *CreateOrderRequest) Send() (interface{}, error) {
	err := r.service.GetAccessToken()
	if err != nil {
		return nil, err
	}

	requestBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type":               "flight-order",
			"flightOffers":       r.flightOffers,
			"travelers":          r.travelers,
			"remarks":            r.remarks,
			"ticketingAgreement": r.ticketingAgreement,
			"contacts":           r.contacts,
		},
	}

	resp, err := r.service.Client.R().
		SetBody(requestBody).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+r.service.AccessToken).
		Post(r.service.BaseUrl + "/v2/booking/flight-orders")

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
