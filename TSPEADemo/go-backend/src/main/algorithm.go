package main

import (
	"math/rand"
	"sort"
)

// Moved all code related to running to the Algorithm Runner
type Algorithm struct {
	population          CandidateSolutions
	mutationProbability int
	populationSize      int
	nrOfGenerations     int
	fitnessThreshold    int
	parentSelectionSize int
	parentPoolSize      int

	// Left out thread and running because this should be in the algoritm runner
	// TODO verify need of overlap between options and algorithm
	// TODO consider that Algorithm contains an OPtions object instead of a own version of each property
	// TODO implement
}

func (algorithm Algorithm) createOffspring(parents CandidateSolutions) CandidateSolutions {
	offspring := make(CandidateSolutions, len(parents))
	for i := 0; i < len(parents); i += 2 {
		parent1 := parents[i]
		parent2 := parents[i + 1]
		children := parent1.recombine(parent2)
		for _, child := range children {
			if (algorithm.shouldBeMutated()) {
				child.mutate()
			}
			offspring = append(offspring, child)
		}
	}
	return offspring
}

func (algorithm Algorithm) shouldBeMutated() bool {
	return rand.Intn(101) <= algorithm.mutationProbability
}

/**
* Selects the survivors by removing the worst candidate solutions from the
* list, so we have the original population size again
*/
// TODO Verify that we are really changing algorithm.population here
// TODO could well be that we need a pointer instead of the current construction
func (algorithm Algorithm) selectSurvivors() {
	// sort population based on fitness
	sort.Sort(algorithm.population)

	// Cut down the population
	algorithm.population = append(algorithm.population, algorithm.population[0:algorithm.populationSize]...)
}

/**
* Select the x best candidate solutions from a randomly selected pool from
* the population
*/
func (algorithm Algorithm) parentSelection() CandidateSolutions {
	// TODO Verify if we need current len of just the fullsize from populationSize
	tempPopulation := make(CandidateSolutions, len(algorithm.population))
	copy(tempPopulation, algorithm.population)

	randomCandidates := make(CandidateSolutions, algorithm.parentPoolSize)
	for i := 0; i < algorithm.parentPoolSize; i++ {
		randomIndex := rand.Intn(len(tempPopulation))
		randomCandidateSolution := tempPopulation[randomIndex]

		// Add candidate to the random candidates
		randomCandidates = append(randomCandidates, randomCandidateSolution)

		// TODO verify
		/*
	   	* delete the candidate from the temp population, so we can't pick
	   	* it again
	   	*/
		tempPopulation = append(tempPopulation[:randomIndex], tempPopulation[randomIndex + 1:]...)
	}

	/* Sort the population so that the best candidates are up front */
	sort.Sort(randomCandidates)

	/*
	* return a list with size parentSelectionSize with the best
	* CandidateSolutions
	*/
	return append(randomCandidates, randomCandidates[0:algorithm.parentSelectionSize]...)
}

func (algorithm Algorithm) initialisation() CandidateSolutions {
	tempPopulation := make(CandidateSolutions, algorithm.populationSize)

	for i := 0; i < algorithm.populationSize; i++ {
		candidateSolution := NewCandidateSolution(getBaseCity(), getRandomizedCities())
		tempPopulation = append(tempPopulation, candidateSolution)
	}
	return tempPopulation
}

// No ternary operator in Go :)
func (algorithm Algorithm) getCurrentBest() {
	if len(algorithm.population) > 0 {
		return algorithm.population[0]
	} else {
		return nil
	}
}


