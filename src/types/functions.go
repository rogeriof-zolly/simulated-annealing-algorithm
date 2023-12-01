package types

import "anneling/src/structures"

type RandomSolution func(*structures.Box) structures.Backpack
type ValidationFunction func(structures.Backpack) int
