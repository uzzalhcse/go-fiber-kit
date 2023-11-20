package main

//
//import (
//	"fmt"
//	"github.com/uzzalhcse/amadeus-go/pkg/amadeus"
//	"github.com/uzzalhcse/amadeus-go/pkg/amadeus/flight"
//)
//
//func main() {
//	apiKey := "wGhv7zl6MfpoxRGA9YEA3w8dgcKqVqov"
//	apiSecret := "4nccwD5I2O5eh9wy"
//	client := amadeus.NewClient(apiKey, apiSecret)
//
//	flightService := flight.NewFlightService(client)
//
//	models, err := flightService.OfferSearch().
//		OriginLocationCode("DEL").
//		DestinationLocationCode("LON").
//		DepartureDate("2023-12-01").
//		ReturnDate("2023-12-05").
//		Adult("2").
//		Max("5").
//		IncludedAirlineCodes("TG").
//		Get()
//
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	fmt.Println("=================Response:", models)
//	offerPrice, err := flightService.OfferPrice().FlightOffers(models.Data[0]).Send()
//	fmt.Println("=================offerPrice:", offerPrice)
//}
