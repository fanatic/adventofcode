package main

import (
	"fmt"
	"io"
	"log"
)

var destinations = map[string]map[string]int{}

func main() {
	cityList := map[string]int{} // set to hold unique cities
	for {
		start := ""
		end := ""
		distance := 0
		_, err := fmt.Scanf("%s to %s = %d\n", &start, &end, &distance)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if destinations[start] == nil {
			destinations[start] = map[string]int{}
		}
		if destinations[end] == nil {
			destinations[end] = map[string]int{}
		}
		destinations[start][end] = distance
		destinations[end][start] = distance
		cityList[start] = 0
		cityList[end] = 0
	}

	// Convert our set to an array
	cities := []string{}
	for city := range cityList {
		cities = append(cities, city)
	}

	// Now we have to brute force all possibilities
	brute([]string{}, cities)
	fmt.Printf("Found best route %d -> %v\n", bestRouteDistance, bestRoute)
}

var bestRoute = []string{}
var bestRouteDistance = 0

func brute(route []string, remainingCities []string) {
	// Check if this city is valid from last one
	if len(route) > 1 && destinations[route[len(route)-2]][route[len(route)-1]] == 0 {
		//fmt.Printf("Dropping %s -> %s route.\n", route[len(route)-2], route[len(route)-1])
		return
	}
	if len(remainingCities) == 0 {
		//fmt.Printf("%v\n", route)
		dist := calcDistance(route)
		if dist != 0 && dist > bestRouteDistance {
			bestRoute = route
			bestRouteDistance = dist
		}
		return
	}

	for i := range remainingCities {
		//fmt.Printf("%s", remainingCities[i][:2])
		newRemainingCities := make([]string, len(remainingCities))
		copy(newRemainingCities, remainingCities)
		newRemainingCities = append(newRemainingCities[:i], newRemainingCities[i+1:]...)

		newRoute := append(route, remainingCities[i])
		brute(newRoute, newRemainingCities)
	}
}

func calcDistance(route []string) int {
	if len(route) < 2 {
		return 0
	}
	distance := 0
	lastCity := route[0]
	debug := ""
	for _, city := range route[1:] {
		cost := destinations[lastCity][city]
		debug += fmt.Sprintf("%s>%s(%d) ", lastCity[:2], city[:2], destinations[lastCity][city])

		if cost == 0 {
			return 0 // Impossible path
		}
		distance += cost
		lastCity = city
	}

	fmt.Printf("%s => %d\n", debug, distance)
	return distance
}
