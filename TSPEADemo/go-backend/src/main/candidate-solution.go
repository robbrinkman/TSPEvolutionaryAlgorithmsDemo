package main

type CandidateSolution struct {
	BaseCity       City
	VisitingCities []City
	Route          []City
	Fitness        float64
	Generation     int64
}

// TODO it seems like passing a pointer here is a more generic solution, don't know why (*CandidateSolution)
func NewCandidateSolution(baseCity City, visitingCities []City) CandidateSolution {
	candidateSolution := CandidateSolution{}
	candidateSolution.BaseCity = baseCity
	candidateSolution.VisitingCities = visitingCities
	candidateSolution.Route = append(candidateSolution.Route, baseCity)
	candidateSolution.Route = append(candidateSolution.Route, visitingCities...)
	candidateSolution.Route = append(candidateSolution.Route, baseCity)
	return candidateSolution
}


// TODO Implement
func (candidateSolution CandidateSolution) recombine(otherParent CandidateSolution) CandidateSolutions {
	return CandidateSolutions{}
}

// TODO Implement
func (candidateSolution CandidateSolution) crossFill(childRoute []City, parentRoute []City, cutIndex int32) {

}

// TODO Implement
func (candidateSolution CandidateSolution) mutate() {

}



// TODO compareTo not implemented here, verify sorting


// TODO optimize with caching on fitness, Memoize? (https://godoc.org/github.com/BenLubar/memoize)
func (candidateSolution CandidateSolution) calculateFitness() int {
	totalDistance := 0
	for i := 0; i < (len(candidateSolution.Route) - 1); i++ {
		city := candidateSolution.Route[i]
		nextCity := candidateSolution.Route[i + 1]
		totalDistance += city.calculateDistance(nextCity)
	}
	return totalDistance
}

type CandidateSolutions []CandidateSolution


// Implementing sort.Interface: https://golang.org/src/sort/sort.go
func (candidateSolutions CandidateSolutions) Len() int {
	return len(candidateSolutions)
}

func (candidateSolutions CandidateSolutions) Less(i int, j int) bool {
	return candidateSolutions[i].calculateFitness() < candidateSolutions[j].calculateFitness()
}

func (candidateSolutions CandidateSolutions) Swap(i int, j int)  {
	candidateSolutions[i], candidateSolutions[j] = candidateSolutions[j], candidateSolutions[i]
}



