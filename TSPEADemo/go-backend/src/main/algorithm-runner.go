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
	fmt.Println("Starting Algorithm")
  algorithm := NewAlgorithm()
	traveler := AlgorithmRunner{Running: true}
	fmt.Println("start")
	traveler.Algorithm = algorithm
	traveler.Algorithm.startAlgorithm()
	return &traveler

}

func (traveler *AlgorithmRunner) stop() {
	traveler.Running = false
}

func (traveler *AlgorithmRunner) getCurrentBest() CandidateSolution {
	return traveler.Algorithm.getCurrentBest()
}
