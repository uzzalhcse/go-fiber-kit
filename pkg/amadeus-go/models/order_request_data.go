// amd/flight/models/flight_offer_data.go
package models

// Contact represents contact information.
type Contact struct {
	AddresseeName struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"addresseeName"`
	CompanyName  string  `json:"companyName"`
	Purpose      string  `json:"purpose"`
	Phones       []Phone `json:"phones"`
	EmailAddress string  `json:"emailAddress"`
	Address      Address `json:"address"`
}

// Address represents address information.
type Address struct {
	Lines       []string `json:"lines"`
	PostalCode  string   `json:"postalCode"`
	CityName    string   `json:"cityName"`
	CountryCode string   `json:"countryCode"`
}

// Phone represents phone information.
type Phone struct {
	DeviceType         string `json:"deviceType"`
	CountryCallingCode string `json:"countryCallingCode"`
	Number             string `json:"number"`
}

// Traveler represents traveler information.
type Traveler struct {
	ID          string     `json:"id"`
	DateOfBirth string     `json:"dateOfBirth"`
	Gender      string     `json:"gender"`
	Contact     Contact    `json:"contact"`
	Name        Name       `json:"name"`
	Documents   []Document `json:"documents"`
}

// Name represents name information.
type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Document represents document information.
type Document struct {
	DocumentType     string `json:"documentType"`
	BirthPlace       string `json:"birthPlace"`
	IssuanceLocation string `json:"issuanceLocation"`
	IssuanceDate     string `json:"issuanceDate"`
	Number           string `json:"number"`
	ExpiryDate       string `json:"expiryDate"`
	IssuanceCountry  string `json:"issuanceCountry"`
	ValidityCountry  string `json:"validityCountry"`
	Nationality      string `json:"nationality"`
	Holder           bool   `json:"holder"`
}

// Remarks represents remarks information.
type Remarks struct {
	General []GeneralRemark `json:"general"`
}

// GeneralRemark represents a general remark.
type GeneralRemark struct {
	SubType string `json:"subType"`
	Text    string `json:"text"`
}

// TicketingAgreement represents ticketing agreement information.
type TicketingAgreement struct {
	Option string `json:"option"`
	Delay  string `json:"delay"`
}
