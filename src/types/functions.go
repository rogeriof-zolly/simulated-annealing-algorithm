package types

import "anneling/src/structures"

type RandomSolution func(bp structures.Backpack, box *structures.Box) structures.Backpack
type ValidationFunction func(structures.Backpack) int
