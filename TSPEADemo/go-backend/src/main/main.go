package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
)

// TODO investigate project structure
// TODO investigate dependency management (no go get?)
// TODO investigate error handling and logging
// TODO investigate unit testing
// TODO investigate naming conventions (methods lower?)

// DONE add rest service for cities


func main() {
	fs := http.FileServer(http.Dir("../frontend/app"))

	router := mux.NewRouter()
	router.HandleFunc("/api/cities", CityIndex)

	router.PathPrefix("/").Handler(fs)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}

func CityIndex(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	cities := Cities{
		City{Name: "Athens", Latitude: 37.975334, Longitude: 23.736151},
		City{Name: "Bern", Latitude: 46.947922, Longitude: 7.444608},
		City{Name: "Berlin", Latitude: 52.519173, Longitude: 13.406091},
		City{Name: "Bratislava", Latitude: 48.146240, Longitude: 17.107262},
		City{Name: "Brussels", Latitude: 50.850342, Longitude: 4.351710},
		City{Name: "Bucharest", Latitude: 44.437710, Longitude: 26.097366},
		City{Name: "Budapest", Latitude: 47.498405, Longitude: 19.040758},
		City{Name: "Copenhagen", Latitude: 55.676098, Longitude: 12.568337},
		City{Name: "Dublin", Latitude: 53.344105, Longitude: -6.267494},
		City{Name: "Helsinki", Latitude: 60.166588, Longitude: 24.943556},
		City{Name: "Lisbon", Latitude: 38.706932, Longitude: -9.135632},
		City{Name: "London", Latitude: 51.508129, Longitude: -0.128005},
		City{Name: "Luxemburg", Latitude: 49.611622, Longitude: 6.131935},
		City{Name: "Madrid", Latitude: 40.416691, Longitude: -3.700345},
		City{Name: "Oslo", Latitude:  59.913868, Longitude: 10.752245},
		City{Name: "Prague", Latitude: 50.087811, Longitude: 14.420460},
		City{Name: "Rome", Latitude: 41.890518, Longitude: 12.494249},
		City{Name: "Sofia", Latitude: 42.696491, Longitude: 23.326012},
		City{Name: "Stockholm", Latitude: 59.328930, Longitude: 18.064911},
		City{Name: "Vienna", Latitude: 48.208176, Longitude: 16.373819},
		City{Name: "Warsaw", Latitude: 52.229675, Longitude: 21.012230}};
	json.NewEncoder(response).Encode(cities)
}
