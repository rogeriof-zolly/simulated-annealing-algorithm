package types

import "anneling/src/structures"

type RandomSolution func(structures.Backpack) structures.Backpack
type ValidationFunction func(structures.Backpack) int
