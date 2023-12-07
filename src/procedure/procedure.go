package procedure

import (
	"anneling/src/structures"
	"anneling/src/types"
	"math"
	"math/rand"
)

// Função base para o algoritmo de simulated annealing
func Process(
	validationFunction types.ValidationFunction,
	randomSolution types.RandomSolution,
	alpha float64,
	maxIterations int,
	initalTemp float64,
	solution structures.Backpack,
	possibleItems structures.Box,
) structures.Backpack {
	bestSolution := solution
	currentIteration := 0
	currentTemperatue := initalTemp

	var (
		neighbor         structures.Backpack
		newPossibleItems structures.Box
		delta            int
	)

	for currentTemperatue > 0 {
		for currentIteration < maxIterations {
			currentIteration += 1
			neighbor, newPossibleItems = randomSolution(solution, possibleItems)
			delta = validationFunction(neighbor) - validationFunction(solution)
			if delta > 0 {
				solution = neighbor
				possibleItems = newPossibleItems
				if validationFunction(neighbor) > validationFunction(bestSolution) {
					bestSolution = neighbor
				}
			} else {
				x := rand.Float64()
				if x < math.Pow(math.E, (-float64(delta)/float64(currentTemperatue))) {
					solution = neighbor
					possibleItems = newPossibleItems
				}
			}
		}
		currentTemperatue = alpha * currentTemperatue
		currentIteration = 0
	}
	solution = bestSolution
	return solution
}
