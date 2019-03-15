package main

import (
	"fmt"
)

func search(street string) {
	result, err := geo.NewSearch().
		Street(street).
		Search()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nsearch by street %s: %s\n", street, result[0].Name)
	}
}

func reverse(latitude float64, longitude float64) {
	result, err := geo.NewSearch().
		Latitude(latitude).
		Longitude(longitude).
		Reverse()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nsearch by latitude %f, longitude: %f: %s\n", latitude, longitude, result[0].Name)
	}
}
