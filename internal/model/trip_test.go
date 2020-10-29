package model

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParseTrip(t *testing.T) {
	trip := Trip{}

	data := []string{
		"ABC",
		"DEF",
		"5",
	}

	trip.parse(data)

	assert.Equal(t, trip.Route, []string{"ABC"})
	assert.Equal(t, trip.Next, data[1])
	assert.Equal(t, trip.Cost, 5.00)
}

func TestSumCostTrip(t *testing.T) {
	trip := Trip{
		Cost: 5,
	}

	trip.sum("15")

	assert.Equal(t, trip.Cost, 20.00)
}

func TestEndOfRouteTrip(t *testing.T) {
	trip := Trip{
		Next:    "DEF",
		Arrival: "DEF",
	}

	assert.Equal(t, trip.endOfRoute(), true)
}

func TestTripToString(t *testing.T) {
	expect := "ABC - DEF > $5.00"
	trip := Trip{
		Route: []string{"ABC", "DEF"},
		Cost:  5,
	}

	assert.Equal(t, trip.ToString(), expect)
}

func TestBuildRoute(t *testing.T) {
	expect := []string{"ABC", "DEF", "GHI", "JKL", "MNO"}

	trip := Trip{
		Departure: "ABC",
		Arrival:   "MNO",
	}

	index := 3

	routes := [][]string{
		[]string{"MNO", "DEF", "75"},
		[]string{"DEF", "GHI", "75"},
		[]string{"GHI", "JKL", "5"},
		[]string{"ABC", "DEF", "10"},
		[]string{"JKL", "MNO", "5"},
		[]string{"PQR", "STU", "40"},
		[]string{"ABC", "STU", "40"},
		[]string{"GHI", "DEF", "75"},
		[]string{"MNO", "PQR", "75"},
	}

	trip.BuildRoute(routes, index)

	assert.Equal(t, trip.Route, expect)
}
