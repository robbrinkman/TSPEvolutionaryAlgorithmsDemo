package main

import "math/rand"

type Traveler struct {
	Running  bool
	Value   float32
}

func NewTraveler() *Traveler {
	rand.Seed(42)
	traveler := Traveler{ Running:true}

	go func() {
		for true == true {
			traveler.Value = rand.Float32()
		}
	}()
	return &traveler
}