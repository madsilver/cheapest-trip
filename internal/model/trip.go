package model

import (
	"fmt"
	"strconv"
	"strings"
)

const DEPARTURE = 0
const ARRIVAL = 1
const COST = 2

// Trip model
type Trip struct {
	Departure string   `json:"departure"`
	Arrival   string   `json:"arrival"`
	Next      string   `json:"-"`
	Route     []string `json:"best_route"`
	Cost      float64  `json:"cost"`
}

// BuildRoute compile the trip route
func (trip *Trip) BuildRoute(routes [][]string, index int) (bool, error) {
	err := trip.parse(routes[index])

	if trip.endOfRoute() {
		trip.Route = append(trip.Route, trip.Arrival)
		return true, err
	}

	for j, r := range routes {
		if strings.EqualFold(trip.Next, r[DEPARTURE]) {
			return trip.BuildRoute(routes, j)
		}
	}

	return false, err
}

func (trip *Trip) parse(data []string) error {
	trip.Route = append(trip.Route, data[DEPARTURE])
	trip.Next = data[ARRIVAL]
	return trip.sum(data[COST])
}

func (trip *Trip) sum(cost string) error {
	v, err := strconv.ParseFloat(cost, 64)
	if err != nil {
		return fmt.Errorf("value '%s' in column cost not allowed", cost)
	}
	trip.Cost += v
	return nil
}

func (trip *Trip) endOfRoute() bool {
	return strings.EqualFold(trip.Next, trip.Arrival)
}

// ToString generates a string from the trip
func (trip *Trip) ToString() string {
	if len(trip.Route) == 0 {
		return "no route found"
	}
	route := strings.Join(trip.Route, " - ")
	return fmt.Sprintf("%s > $%.2f", route, trip.Cost)
}
