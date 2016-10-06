package main

import (
	"math/rand"
	"log"
	"strings"
)

type CandidateSolution struct {
	BaseCity       City
	VisitingCities []City
	Route          []City        `json:"route"`
	Fitness        float64  `json:"fitness"`
	Generation     int
}

// TODO it seems like passing a pointer here is a more generic solution, don't know why (*CandidateSolution)
func NewCandidateSolution(baseCity City, visitingCities []City) CandidateSolution {
	candidateSolution := CandidateSolution{}
	candidateSolution.BaseCity = baseCity
	candidateSolution.VisitingCities = visitingCities
	candidateSolution.Route = append(candidateSolution.Route, baseCity)
	candidateSolution.Route = append(candidateSolution.Route, visitingCities...)
	candidateSolution.Route = append(candidateSolution.Route, baseCity)
	candidateSolution.calculateFitness()
	return candidateSolution
}


// TODO Implement
func (candidateSolution *CandidateSolution) recombine(otherParent CandidateSolution) CandidateSolutions {

	/* get routes of both parents */
	parentRoute1 := candidateSolution.VisitingCities
	parentRoute2 := otherParent.VisitingCities

	/* randomize cutIndex for "cross-and-fill point" */
	cutIndex :=  int32(rand.Intn(len(parentRoute1)))

	/* initialize the routes for the children */
	childRoute1 := make(Cities, len(parentRoute1))
	childRoute2 := make(Cities, len(parentRoute1))



	/* get the first part of both parent routes using the cut index */
	partRoute1 := parentRoute1[0:cutIndex]
	partRoute2 := parentRoute2[0:cutIndex]

	/* copy the first part of the parents cut into the children */
	copy(childRoute1, partRoute1)
	copy(childRoute2, partRoute2)


	/*
	 * Now, the "difficult part". Check the rest of the route in the
	 * crossing parent and add the cities that are not yet in the child (in
	 * the order of the route of the crossing parent)
	 */
	candidateSolution.crossFill(childRoute1, parentRoute2, cutIndex)
	candidateSolution.crossFill(childRoute2, parentRoute1, cutIndex)

	/* create new children using the new children routes */
	child1 := NewCandidateSolution(getBaseCity(), childRoute1);
	child2 := NewCandidateSolution(getBaseCity(), childRoute2);

	/* put the children in a list and return it */

	return CandidateSolutions{child1, child2}
}

func (candidateSolution CandidateSolution) printRoute() {
	cityNames:= make([]string, len(candidateSolution.VisitingCities))
	for i, city := range candidateSolution.VisitingCities {
		cityNames[i] = city.Name
	}
	log.Printf("Route -> %s", strings.Join(cityNames, " -> "))
}

/**
* Check the rest of the route in the crossing parent and add the cities
* that are not yet in the child (in the order of the route of the crossing
* parent)
*/

// TODO I guess childRoute should be a pointer
func (candidateSolution *CandidateSolution) crossFill(childRoute Cities, parentRoute []City, cutIndex int32) {
	/*
	 * traverse the parent route from the cut index on and add every city
	 * not yet in the child to the child
	 */
	fillIndex := cutIndex

	for i := cutIndex; i < int32(len(parentRoute)); i++ {
		nextCityOnRoute := parentRoute[i]
		if (!childRoute.contains(nextCityOnRoute)) {
			childRoute[fillIndex] = nextCityOnRoute
			fillIndex++
		}
	}

	/*
	 * traverse the parent route from the start of the route and add every
	 * city not yet in the child to the child
	 */

	for i := 0; i < int(cutIndex); i++ {
		nextCityOnRoute := parentRoute[i]
		if (!childRoute.contains(nextCityOnRoute)) {
			childRoute[fillIndex] = nextCityOnRoute
			fillIndex++
		}
	}
}

func (candidateSolution *CandidateSolution) mutate() {

	/* randomly select two indices in the route */
	indexFirstCity := int32(rand.Intn(len(candidateSolution.VisitingCities)))
	indexSecondCity := int32(rand.Intn(len(candidateSolution.VisitingCities)))

	/* Make sure they are different */
	for (indexFirstCity == indexSecondCity) {
		indexSecondCity = int32(rand.Intn(len(candidateSolution.VisitingCities)))
	}

	/* Changer! */
	candidateSolution.VisitingCities[indexFirstCity], candidateSolution.VisitingCities[indexSecondCity] = candidateSolution.VisitingCities[indexSecondCity], candidateSolution.VisitingCities[indexFirstCity]


	// fitness changes. Since we are doing caching:
	candidateSolution.calculateFitness()
}

func (candidateSolution *CandidateSolution) getFitness() float64 {
	if (candidateSolution.Fitness == 0) {
		candidateSolution.calculateFitness()
	}
	return candidateSolution.Fitness
}

func (candidateSolution *CandidateSolution) calculateFitness() {
	totalDistance := float64(0)
	for i := 0; i < (len(candidateSolution.Route) - 1); i++ {
		city := candidateSolution.Route[i]
		nextCity := candidateSolution.Route[i + 1]
		totalDistance += city.calculateDistance(nextCity)
	}
	candidateSolution.Fitness = totalDistance
}

type CandidateSolutions []CandidateSolution


// Implementing sort.Interface: https://golang.org/src/sort/sort.go
func (candidateSolutions CandidateSolutions) Len() int {
	return len(candidateSolutions)
}

func (candidateSolutions CandidateSolutions) Less(i int, j int) bool {
	return candidateSolutions[i].getFitness() < candidateSolutions[j].getFitness()
}

func (candidateSolutions CandidateSolutions) Swap(i int, j int) {
	candidateSolutions[i], candidateSolutions[j] = candidateSolutions[j], candidateSolutions[i]
}
