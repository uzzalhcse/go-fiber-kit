package flight

import "github.com/uzzalhcse/amadeus-go/pkg/amadeus"

type Service struct {
	client *amadeus.Client
}

func NewFlightService(client *amadeus.Client) *Service {
	return &Service{client: client}
}

func (s *Service) OfferSearch() *SearchRequest {
	return &SearchRequest{service: s}
}

func (s *Service) OfferPrice() *PricingRequest {
	return &PricingRequest{service: s}
}
