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

	// Set to running
	traveler := Traveler{Running: true}
	fmt.Println("start")
	go func() {
		for traveler.Running == true {
			traveler.Value = rand.Float32()
			fmt.Println(traveler.Value)
			time.Sleep(1000*time.Millisecond)
			//if (traveler.Value<0.00000001) {
			//	traveler.Running = false
			//}
			//time.Sleep(100 * time.Millisecond)
		}
	}()
	return &traveler

}

func (traveler *Traveler) stop() {
	traveler.Running = false
}