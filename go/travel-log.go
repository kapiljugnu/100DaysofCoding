package main

import (
	"fmt"
)

type TravelLog struct {
	country string
	visits  int
	cities  []string
}

var travel_log []TravelLog = []TravelLog{
	{
		country: "France",
		visits:  12,
		cities:  []string{"Paris", "Lille", "Dijon"},
	},
	{
		country: "Germany",
		visits:  5,
		cities:  []string{"Berlin", "Hamburg", "Stuttgart"},
	},
}

func add_new_country(country_visited string, times_visited int, cities_visited []string) {
	new_country := TravelLog{
		country: country_visited,
		visits:  times_visited,
		cities:  cities_visited,
	}
	travel_log = append(travel_log, new_country)
}

func main() {
	add_new_country("Russia", 2, []string{"Moscow", "Saint Petersburg"})
	fmt.Println(travel_log)
}
