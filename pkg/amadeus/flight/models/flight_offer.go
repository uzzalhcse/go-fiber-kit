package models

type FlightOfferData struct {
	Meta         Meta          `json:"meta"`
	Data         []FlightOffer `json:"data"`
	Dictionaries Dictionaries  `json:"dictionaries"`
}

type Meta struct {
	Count int `json:"count"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type Dictionaries struct {
	Locations  map[string]Location `json:"locations"`
	Aircraft   map[string]string   `json:"aircraft"`
	Currencies map[string]string   `json:"currencies"`
	Carriers   map[string]string   `json:"carriers"`
}

type Location struct {
	CityCode    string `json:"cityCode"`
	CountryCode string `json:"countryCode"`
}

type FlightOffer struct {
	Type                     string            `json:"type"`
	ID                       string            `json:"id"`
	Source                   string            `json:"source"`
	InstantTicketingRequired bool              `json:"instantTicketingRequired"`
	NonHomogeneous           bool              `json:"nonHomogeneous"`
	OneWay                   bool              `json:"oneWay"`
	LastTicketingDate        string            `json:"lastTicketingDate"`
	LastTicketingDateTime    string            `json:"lastTicketingDateTime"`
	NumberOfBookableSeats    int               `json:"numberOfBookableSeats"`
	Itineraries              []Itinerary       `json:"itineraries"`
	Price                    Price             `json:"price"`
	PricingOptions           PricingOptions    `json:"pricingOptions"`
	ValidatingAirlineCodes   []string          `json:"validatingAirlineCodes"`
	TravelerPricings         []TravelerPricing `json:"travelerPricings"`
}

type Itinerary struct {
	Duration string    `json:"duration"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Departure       Departure `json:"departure"`
	Arrival         Arrival   `json:"arrival"`
	CarrierCode     string    `json:"carrierCode"`
	Number          string    `json:"number"`
	Aircraft        Aircraft  `json:"aircraft"`
	Operating       Operating `json:"operating"`
	Duration        string    `json:"duration"`
	ID              string    `json:"id"`
	NumberOfStops   int       `json:"numberOfStops"`
	BlacklistedInEU bool      `json:"blacklistedInEU"`
}

type Departure struct {
	IATACode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Arrival struct {
	IATACode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Aircraft struct {
	Code string `json:"code"`
}

type Operating struct {
	CarrierCode string `json:"carrierCode"`
}

type Price struct {
	Currency           string              `json:"currency"`
	Total              string              `json:"total"`
	Base               string              `json:"base"`
	Fees               []Fee               `json:"fees"`
	GrandTotal         string              `json:"grandTotal"`
	AdditionalServices []AdditionalService `json:"additionalServices"`
}

type Fee struct {
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type AdditionalService struct {
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type PricingOptions struct {
	FareType                []string `json:"fareType"`
	IncludedCheckedBagsOnly bool     `json:"includedCheckedBagsOnly"`
}

type TravelerPricing struct {
	TravelerID           string                 `json:"travelerId"`
	FareOption           string                 `json:"fareOption"`
	TravelerType         string                 `json:"travelerType"`
	Price                Price                  `json:"price"`
	FareDetailsBySegment []FareDetailsBySegment `json:"fareDetailsBySegment"`
}

type FareDetailsBySegment struct {
	SegmentID           string              `json:"segmentId"`
	Cabin               string              `json:"cabin"`
	FareBasis           string              `json:"fareBasis"`
	BrandedFare         string              `json:"brandedFare"`
	Class               string              `json:"class"`
	IncludedCheckedBags IncludedCheckedBags `json:"includedCheckedBags"`
}

type IncludedCheckedBags struct {
	Quantity int `json:"quantity"`
}
