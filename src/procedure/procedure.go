package procedure

import (
	"anneling/src/structures"
	"anneling/src/types"
	"math"
	"math/rand"
)

func Process(
	validationFunction types.ValidationFunction,
	randomSolution types.RandomSolution,
	alpha int,
	maxIterations int,
	initalTemp int,
	solution structures.Backpack,
	possibleItems structures.Box,
) structures.Backpack {
	bestSolution := solution
	currentIteration := 0
	currentTemperatue := initalTemp

	for currentTemperatue > 0 {
		for currentIteration < maxIterations {
			currentIteration += 1
			neighbor := randomSolution(&possibleItems)
			delta := validationFunction(neighbor) - validationFunction(solution)
			if delta > 0 {
				solution = neighbor
				if validationFunction(neighbor) > validationFunction(bestSolution) {
					bestSolution = neighbor
				}
			} else {
				x := rand.Float64()
				if x < math.Pow(math.E, (-float64(delta)/float64(currentTemperatue))) {
					solution = neighbor
				}
			}
		}
		currentTemperatue = alpha * currentTemperatue
		currentIteration = 0
	}
	solution = bestSolution
	return solution
}
