package console

import (
	"fmt"

	"github.com/silver/cheapest-trip/internal/core/search"
)

// Run Console
func Run() {
	var scan string

	for {
		fmt.Print("please enter the route: ")
		fmt.Scan(&scan)

		points, err := search.Parse(scan)

		if err != nil {
			fmt.Println(err)
			continue
		}

		trip, err := search.FindCheaper(points)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("best route: " + trip.ToString())
	}
}
