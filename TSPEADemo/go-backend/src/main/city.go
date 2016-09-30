package main

type City struct {
	Name      string
	Latitude  float32
	Longitude float32
}

type Cities []City


func (city City) calculateDistance(otherCity City) int {
	return 10
}