package types

import "anneling/src/structures"

// Tipagem das funções para poder passá-las como parâmetro para o processamento
type RandomSolution func(bp structures.Backpack, box structures.Box) (structures.Backpack, structures.Box)
type ValidationFunction func(structures.Backpack) int
