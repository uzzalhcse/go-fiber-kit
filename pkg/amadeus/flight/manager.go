package flight

import (
	"github.com/uzzalhcse/amadeus-go/pkg/amadeus"
)

type Service struct {
	Client *amadeus.Client
}

func NewFlightService(client *amadeus.Client) *Service {
	return &Service{Client: client}
}

func (s *Service) OfferSearch() *SearchRequest {
	return &SearchRequest{Service: s}
}

func (s *Service) OfferPrice() *PricingRequest {
	return &PricingRequest{Service: s}
}
