package main

import (
	"fmt"
	"github.com/joaosoft/geo-location"
)

var geo, _ = geolocation.NewGeoLocation()

func main() {

	// document create
	fmt.Println(":: SEARCH BY: STREET")
	search("rua particular de monsanto")

	fmt.Println(":: REVERSE BY: LATITUDE/LONGITUDE")
	reverse(41.1718238, -8.6186277)
}
