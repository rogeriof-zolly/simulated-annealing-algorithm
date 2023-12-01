package functions

import "anneling/src/structures"

func ValidationFunction(bp structures.Backpack) int {
	sum := 0

	for _, item := range bp.Items {
		sum += item.Utility
	}

	return sum
}
