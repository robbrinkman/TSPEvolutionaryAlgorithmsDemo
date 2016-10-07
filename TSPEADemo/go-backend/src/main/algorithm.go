package main

import (
	"math/rand"
	"sort"
	"log"
	"time"
)

type Algorithm struct {
	currentBest         CandidateSolution
	population          CandidateSolutions
	mutationProbability int
	populationSize      int
	nrOfGenerations     int
	generations         int
	fitnessThreshold    float64
	parentSelectionSize int
	parentPoolSize      int
	running             bool
}

func NewAlgorithm() Algorithm {
	algorithm := Algorithm{}
	algorithm.mutationProbability = 25
	algorithm.populationSize = 50000
	algorithm.nrOfGenerations = 5000
	algorithm.fitnessThreshold = 6000
	algorithm.parentSelectionSize = 6000
	algorithm.parentPoolSize = 10000
	algorithm.running = false
	return algorithm
}

func (algorithm *Algorithm) stop() {
	algorithm.running = false
}

func (algorithm *Algorithm) start() {

	algorithm.running = true

	algorithm.initialisation()

	go func() {

		for algorithm.generations = 0; algorithm.generations != algorithm.nrOfGenerations  && algorithm.population[0].GetFitness() > algorithm.fitnessThreshold && algorithm.running;
		{
			parents := algorithm.parentSelection()

			offspring := algorithm.createOffspring(parents)

			algorithm.population = append(algorithm.population, offspring...)

			algorithm.determineCurrentBest()
			algorithm.selectSurvivors()

			algorithm.generations++

			log.Printf("Generation: %d", algorithm.generations)
			log.Printf("Current best: %f", algorithm.population[0].GetFitness())
		}
	}()
}

func (algorithm *Algorithm) createOffspring(parents CandidateSolutions) CandidateSolutions {
	//defer timeTrack(time.Now(), "createOffspring")
	offspring := make(CandidateSolutions, len(parents))
	offSpringIndex := 0
	for i := 0; i < len(parents); i += 2 {
		parent1 := parents[i]
		parent2 := parents[i + 1]
		children := parent1.recombine(parent2)
		for _, child := range children {
			if (algorithm.shouldBeMutated()) {
				child.mutate()
			}
			offspring[offSpringIndex] = child
			offSpringIndex++
		}
	}
	return offspring
}

func (algorithm *Algorithm) shouldBeMutated() bool {
	return rand.Intn(101) <= algorithm.mutationProbability
}

/**
* Selects the survivors by removing the worst candidate solutions from the
* list, so we have the original population size again
*/
func (algorithm *Algorithm) selectSurvivors() {
	//defer timeTrack(time.Now(), "selectSurvivors")
	algorithm.population = algorithm.population[0:algorithm.populationSize]
}

/**
* Select the x best candidate solutions from a randomly selected pool from
* the population
*/
func (algorithm *Algorithm) parentSelection() CandidateSolutions {
	//defer timeTrack(time.Now(), "parentSelection")

	tempPopulation := make(CandidateSolutions, algorithm.populationSize)
	copy(tempPopulation, algorithm.population)

	randomCandidates := make(CandidateSolutions, algorithm.parentPoolSize)

	for i := 0; i < algorithm.parentPoolSize; i++ {
		randomIndex := rand.Intn(len(tempPopulation))
		randomCandidateSolution := tempPopulation[randomIndex]

		// Add candidate to the random candidates
		randomCandidates[i] = randomCandidateSolution

		/*
	   	* delete the candidate from the temp population, so we can't pick
	   	* it again
	   	*/

		tempPopulation[randomIndex] = tempPopulation[len(tempPopulation)-1] // Replace it with the last one.
		tempPopulation = tempPopulation[:len(tempPopulation)-1]
	}


	/* Sort the population so that the best candidates are up front */
	sort.Sort(randomCandidates)

	/*
	* return a list with size parentSelectionSize with the best
	* CandidateSolutions
	*/
	result := randomCandidates[0:algorithm.parentSelectionSize]

	return result
}

func (algorithm *Algorithm) initialisation() {
	algorithm.population = make(CandidateSolutions, algorithm.populationSize)
	for i := 0; i < algorithm.populationSize; i++ {
		candidateSolution := NewCandidateSolution(getBaseCity(), getRandomizedCities())
		algorithm.population[i] = candidateSolution
	}
	algorithm.determineCurrentBest()
}

func (algorithm *Algorithm) determineCurrentBest() {
	//defer timeTrack(time.Now(), "determineCurrentBest")
	if (len(algorithm.population)> 0) {
		sort.Sort(algorithm.population)
		algorithm.currentBest = algorithm.population[0]
	}
}

func (algorithm *Algorithm) getCurrentBest() CandidateSolution {
	if len(algorithm.population) > 0 {
		currentBest := algorithm.currentBest
		currentBest.Generation = algorithm.generations
		return currentBest
	} else {
		// TODO find nice solution to return something if no population or throw exception
		return CandidateSolution{}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
