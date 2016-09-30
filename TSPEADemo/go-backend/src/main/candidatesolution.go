package main

type CandidateSolution struct {
	BaseCity       City
	VisitingCities []City
	Route          []City
	Fitness        float64
	Generation     int64
}

type CandidateSolutions []CandidateSolution

// TODO implement alternative for constructor


// TODO optimize with caching on fitness, Memoize?
func (candidateSolution CandidateSolution) calculateFitness() float64 {
	totalDistance := 0
	for i := 0; i < (len(candidateSolution.Route) - 1); i++ {
		city := candidateSolution.Route[i]
		nextCity := candidateSolution.Route[i + 1]
		totalDistance += city.calculateDistance(nextCity)
	}
	return totalDistance
}
