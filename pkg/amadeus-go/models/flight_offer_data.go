// amd/flight/models/flight_offer_data.go
package models

import "encoding/json"

type DateTime struct {
	At string `json:"at"`
}

type Departure struct {
	At       string `json:"at"`
	IataCode string `json:"iataCode"`
	Terminal string `json:"terminal,omitempty"`
}

type Arrival struct {
	At       string `json:"at"`
	IataCode string `json:"iataCode"`
	Terminal string `json:"terminal,omitempty"`
}
type Aircraft struct {
	Code string `json:"code"`
}

type Operating struct {
	CarrierCode string `json:"carrierCode"`
}

type Segment struct {
	ID              string    `json:"id"`
	Aircraft        Aircraft  `json:"aircraft"`
	Arrival         Arrival   `json:"arrival"`
	BlacklistedInEU bool      `json:"blacklistedInEU"`
	CarrierCode     string    `json:"carrierCode"`
	Departure       Departure `json:"departure"`
	Duration        string    `json:"duration"`
	Number          string    `json:"number"`
	NumberOfStops   int       `json:"numberOfStops"`
	Operating       Operating `json:"operating"`
}

type Itinerary struct {
	Duration string    `json:"duration"`
	Segments []Segment `json:"segments"`
}

type Fee struct {
	Amount json.Number `json:"amount"`
	Type   string      `json:"type"`
}

type Price struct {
	Base       json.Number `json:"base"`
	Currency   string      `json:"currency"`
	Fees       []Fee       `json:"fees"`
	GrandTotal json.Number `json:"grandTotal"`
	Total      json.Number `json:"total"`
}

type FareDetailsBySegment struct {
	Cabin               string `json:"cabin"`
	Class               string `json:"class"`
	FareBasis           string `json:"fareBasis"`
	IncludedCheckedBags struct {
		Quantity int `json:"quantity"`
	} `json:"includedCheckedBags"`
	SegmentID string `json:"segmentId"`
}

type TravelerPricing struct {
	FareDetailsBySegment []FareDetailsBySegment `json:"fareDetailsBySegment"`
	FareOption           string                 `json:"fareOption"`
	Price                struct {
		Currency string      `json:"currency"`
		Total    json.Number `json:"total"`
		Base     json.Number `json:"base"`
	} `json:"price"`
	TravelerId   string `json:"travelerId"`
	TravelerType string `json:"travelerType"`
}

type FlightOffer struct {
	ID                       string      `json:"id"`
	InstantTicketingRequired bool        `json:"instantTicketingRequired"`
	Itineraries              []Itinerary `json:"itineraries"`
	LastTicketingDate        string      `json:"lastTicketingDate"`
	LastTicketingDateTime    string      `json:"lastTicketingDateTime"`
	NonHomogeneous           bool        `json:"nonHomogeneous"`
	NumberOfBookableSeats    int         `json:"numberOfBookableSeats"`
	OneWay                   bool        `json:"oneWay"`
	Price                    Price       `json:"price"`
	PricingOptions           struct {
		FareType                []string `json:"fareType"`
		IncludedCheckedBagsOnly bool     `json:"includedCheckedBagsOnly"`
	} `json:"pricingOptions"`
	Source                 string            `json:"source"`
	TravelerPricings       []TravelerPricing `json:"travelerPricings"`
	Type                   string            `json:"type"`
	ValidatingAirlineCodes []string          `json:"validatingAirlineCodes"`
}

type AircraftDictionary map[string]string

type CarrierDictionary map[string]string

type CurrencyDictionary map[string]string

type LocationDictionary map[string]struct {
	CityCode    string `json:"cityCode"`
	CountryCode string `json:"countryCode"`
}

type Dictionaries struct {
	Aircraft   AircraftDictionary `json:"aircraft"`
	Carriers   CarrierDictionary  `json:"carriers"`
	Currencies CurrencyDictionary `json:"currencies"`
	Locations  LocationDictionary `json:"locations"`
}

type Meta struct {
	Count int `json:"count"`
}

type FlightOfferData struct {
	Data         []FlightOffer `json:"data"`
	Dictionaries Dictionaries  `json:"dictionaries"`
	Meta         Meta          `json:"meta"`
}
