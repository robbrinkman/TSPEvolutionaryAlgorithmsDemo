package main

type Traveler struct {
	Running      bool
	Value  float32
}

func NewTraveler() Traveler  {
	traveler := Traveler{Running:true}
	return traveler
}