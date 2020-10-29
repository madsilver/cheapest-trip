package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/silver/cheapest-trip/internal/core/search"
	"github.com/silver/cheapest-trip/internal/core/storage"
	"github.com/silver/cheapest-trip/internal/model"
)

func get(res http.ResponseWriter, req *http.Request) {
	route := req.URL.Query().Get("route")

	points, err := search.Parse(route)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	trip, err := search.FindCheaper(points)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonBody, _ := json.Marshal(trip)
	res.Write(jsonBody)
}

func post(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	trip := model.Trip{}
	err := decoder.Decode(&trip)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	route := []string{
		trip.Departure,
		trip.Arrival,
		fmt.Sprintf("%.2f", trip.Cost),
	}

	if !validate(route) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	routes, _ := storage.LoadCsv(os.Args)
	routes = append(routes, route)

	err = storage.WriteCsv(os.Args, routes)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func routerHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		get(res, req)
	case "POST":
		post(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func validate(data []string) bool {
	valid := true
	for _, v := range data {
		if v == "" {
			valid = false
			break
		}
	}
	return valid
}

// Run Server
func Run() {
	http.HandleFunc("/routes", routerHandler)

	fmt.Println(http.ListenAndServe(":9000", nil))
}
