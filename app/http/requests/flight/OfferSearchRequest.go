package flightrequests

import "github.com/uzzalhcse/amadeus-go/pkg/validator"

type DateTimeRange struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type OriginDestination struct {
	ID                      string        `json:"id"`
	OriginLocationCode      string        `json:"originLocationCode"`
	DestinationLocationCode string        `json:"destinationLocationCode"`
	DepartureDateTimeRange  DateTimeRange `json:"departureDateTimeRange"`
}

type FareOption struct {
	TravelerType string   `json:"travelerType"`
	FareOptions  []string `json:"fareOptions"`
}

type Traveler struct {
	ID           string   `json:"id"`
	TravelerType string   `json:"travelerType"`
	FareOptions  []string `json:"fareOptions"`
}

type CabinRestriction struct {
	Cabin                string   `json:"cabin"`
	Coverage             string   `json:"coverage"`
	OriginDestinationIds []string `json:"originDestinationIds"`
}

type CarrierRestrictions struct {
	ExcludedCarrierCodes []string `json:"excludedCarrierCodes"`
}

type FlightFilters struct {
	CabinRestrictions   []CabinRestriction  `json:"cabinRestrictions"`
	CarrierRestrictions CarrierRestrictions `json:"carrierRestrictions"`
}

type SearchCriteria struct {
	MaxFlightOffers int           `json:"maxFlightOffers"`
	FlightFilters   FlightFilters `json:"flightFilters"`
}

type FlightSearchRequest struct {
	CurrencyCode       string              `json:"currencyCode"`
	OriginDestinations []OriginDestination `json:"originDestinations"`
	Travelers          []Traveler          `json:"travelers"`
	Sources            []string            `json:"sources"`
	SearchCriteria     SearchCriteria      `json:"searchCriteria"`
	*validator.Request
}
