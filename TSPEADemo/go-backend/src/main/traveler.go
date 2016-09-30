package main

import (
	"math/rand"
	"fmt"
	"time"
)

type Traveler struct {
	Running bool
	Value   float32
}

// Constructor should use options to configure traveling
//func NewTraveler() Traveler {
//	return Traveler{Running:false}
//}

// Start using a pointer so we can run a go routine on the actual travelr
func startTraveler() *Traveler {

	// Create new traveller and set to running
	// TODO Should accept the config from the client side

	traveler := Traveler{Running: true}
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

func (traveler *Traveler) stop() {
	traveler.Running = false
}