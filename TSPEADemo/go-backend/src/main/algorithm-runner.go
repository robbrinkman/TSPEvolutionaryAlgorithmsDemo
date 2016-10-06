package main

import (
	"fmt"
)

type AlgorithmRunner struct {
	Running bool
	Algorithm Algorithm
}

// TODO implement multi threading here

// Start using a pointer so we can run a go routine on the actual travelr
func startAlgorithmRunner() *AlgorithmRunner {

	// Create new traveller and set to running
	// TODO Should accept the config from the client side
	// TODO Refactor
	fmt.Println("Starting Algorithm")
  algorithm := NewAlgorithm()
	traveler := AlgorithmRunner{Running: true}
	fmt.Println("start")
	traveler.Algorithm = algorithm
	traveler.Algorithm.start()
	return &traveler

}

func (traveler *AlgorithmRunner) stop() {
	traveler.Running = false
	traveler.Algorithm.stop()
}

func (traveler *AlgorithmRunner) getCurrentBest() CandidateSolution {
	return traveler.Algorithm.getCurrentBest()
}
