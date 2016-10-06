package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http/pprof"
)

// TODO investigate project structure
// TODO investigate dependency management (no go get?)
// TODO investigate error handling and logging
// TODO investigate unit testing
// TODO investigate naming conventions (methods lower?)

// TODO implement runner using Channels : http://guzalexander.com/2013/12/06/golang-channels-tutorial.html

var algorithmRunner *AlgorithmRunner

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cities", ListCities)

	// TODO reconsider api structure when all works
	/* TODO suggest /api/travel
		POST will start traveling
		GET will return current and finished (true|false) until finished
		DELETE will stop
	 */

	router.HandleFunc("/api/currentBest", CurrentBest)
	router.HandleFunc("/api/latestBest", LatestBest)
	router.HandleFunc("/api/stillRunning", StillRunning)
	router.HandleFunc("/api/startAlgorithm", StartAlgorithm)
	router.HandleFunc("/api/stopAlgorithm", StopAlgorithm)

	AttachProfiler(router)
	
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/app")))

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

func ListCities(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	cities := getAllCities()
	json.NewEncoder(response).Encode(cities)
}

func CurrentBest(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(algorithmRunner.getCurrentBest())
}

func LatestBest(response http.ResponseWriter, request *http.Request) {
	// TODO: This is just testing REST and FE
	CurrentBest(response, request)
}
func StillRunning(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(algorithmRunner != nil && algorithmRunner.Running)
}

func StartAlgorithm(response http.ResponseWriter, request *http.Request) {
	if (algorithmRunner != nil && algorithmRunner.Running) {
		log.Println("Trying to start an running Traveler, skipping")
	} else {
		algorithmRunner = startAlgorithmRunner()
	}
}

func StopAlgorithm(response http.ResponseWriter, request *http.Request) {
	if (algorithmRunner == nil || !algorithmRunner.Running) {
		log.Println("Trying to stop an non-running Traveler, skipping")
	} else {
		algorithmRunner.stop()
	}
}
