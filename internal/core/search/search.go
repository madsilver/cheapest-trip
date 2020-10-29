package search

import (
	"errors"
	"os"
	"strings"

	"github.com/silver/cheapest-trip/internal/core/storage"
	"github.com/silver/cheapest-trip/internal/model"
)

// FindCheaper finds the cheaper route
func FindCheaper(points []string) (model.Trip, error) {
	routes, _ := storage.LoadCsv(os.Args)
	trips := []model.Trip{}
	var err error

	for index, route := range routes {
		if !isDeparture(route[0], points[0]) {
			continue
		}

		trip := trip(points)

		find, err := trip.BuildRoute(routes, index)
		if err != nil {
			return trip, err
		}

		if find {
			trips = append(trips, trip)
		}
	}

	if len(trips) == 0 {
		return trip(points), err
	}

	return cheaper(&trips), err
}

// Parse parse the value provided by the interface
func Parse(value string) (c []string, e error) {
	c = strings.Split(value, "-")
	if len(c) < 2 {
		e = errors.New("incorrect format of the route given")
	}
	return
}

func isDeparture(a string, b string) bool {
	return strings.EqualFold(a, b)
}

func trip(points []string) model.Trip {
	return model.Trip{
		Departure: strings.ToUpper(points[0]),
		Arrival:   strings.ToUpper(points[1]),
		Cost:      0,
		Route:     make([]string, 0),
	}
}

func cheaper(trips *[]model.Trip) (trip model.Trip) {
	for i, t := range *trips {
		if i == 0 {
			trip = t
			continue
		}
		if trip.Cost > t.Cost {
			trip = t
		}
	}
	return
}
