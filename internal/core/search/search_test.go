package search

import (
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/silver/cheapest-trip/internal/model"
)

func TestParseInput(t *testing.T) {
	expect := []string{"ABC", "DEF"}
	parse, err := Parse("ABC-DEF")
	assert.Equal(t, err, nil)
	assert.Equal(t, parse, expect)
}

func TestErrorParseInput(t *testing.T) {
	_, err := Parse("ABC")
	assert.Equal(t, err.Error(), "incorrect format of the route given")
}

func TestCreateTrip(t *testing.T) {
	points := []string{"ABC", "DEF"}
	expect := model.Trip{
		Departure: strings.ToUpper(points[0]),
		Arrival:   strings.ToUpper(points[1]),
		Cost:      0,
		Route:     make([]string, 0),
	}

	trip := trip(points)
	assert.Equal(t, trip, expect)
}

func TestGetCheaperInSlice(t *testing.T) {
	trips := []model.Trip{
		model.Trip{Cost: 20},
		model.Trip{Cost: 12},
		model.Trip{Cost: 75},
	}

	trip := cheaper(&trips)

	assert.Equal(t, trip.Cost, 12.)
}
