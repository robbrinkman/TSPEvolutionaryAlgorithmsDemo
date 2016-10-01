package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Algorithm struct {
	population CandidateSolutions
	mutationProbability int
	populationSize	int
	nrOfGenerations int
	fitnessThreshold int
	parentSelectionSize int
	parentPoolSize int

	// Left out thread and running because this should be in the algoritm runner
	// TODO verify need of overlap between options and algorithm
	// TODO implement
}

type AlgorithmOptions struct {
	mutationProbability int
	populationSize      int
	nrOfGenerations     int
	fitnessThreshold    int
	parentSelectionSize int
	parentPoolSize      int
	algorithmStyle      string
}

type AlgorithmRunner struct {
	Running bool
	Value   float32
}

// Start using a pointer so we can run a go routine on the actual travelr
func startAlgorithmRunner() *AlgorithmRunner {

	// Create new traveller and set to running
	// TODO Should accept the config from the client side

	traveler := AlgorithmRunner{Running: true}
	fmt.Println("start")
	go func() {
		// TODO actually do the genetic travelling here
		for traveler.Running == true {
			traveler.Value = rand.Float32()
			fmt.Println(traveler.Value)

			// Go is to fast so we sleep :)
			time.Sleep(1000*time.Millisecond)
		}
	}()
	return &traveler

}

func (traveler *AlgorithmRunner) stop() {
	traveler.Running = false
}