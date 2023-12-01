package functions

import (
	"anneling/src/structures"
	"math/rand"
)

func RandomSolution(bp structures.Backpack, box *structures.Box) structures.Backpack {
	var (
		boxPosition      int
		backpackPosition int
	)
	if len(box.Items) > 1 {
		boxPosition = rand.Intn(len(box.Items) - 1)
	} else {
		boxPosition = 0
	}

	if len(bp.GetItems()) > 1 {
		backpackPosition = rand.Intn(len(bp.GetItems()) - 1)
	} else {
		backpackPosition = 0
	}

	bp.Items[backpackPosition], box.Items[boxPosition] = box.Items[boxPosition], bp.Items[backpackPosition]

	return bp
}
