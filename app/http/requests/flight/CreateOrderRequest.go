package flightrequests

import (
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus-go/models"
	"github.com/uzzalhcse/amadeus-go/pkg/validator"
)

// CreateOrderRequest represents the request structure for creating an order.
type CreateOrderRequest struct {
	FlightOfferData    []interface{}             `json:"flightOfferData" validate:"required"`
	Travelers          []models.Traveler         `json:"travelers"`
	Remarks            models.Remarks            `json:"remarks"`
	TicketingAgreement models.TicketingAgreement `json:"ticketingAgreement"`
	Contacts           []models.Contact          `json:"contacts"`
	*validator.Request                           // Embedding validator.Request for validation purposes
}
