package procedure

import (
	"anneling/src/functions"
	"anneling/src/structures"
	"fmt"
	"testing"
)

func TestProcedure(t *testing.T) {
	t.Run("Should get the hightest values in a short box", func(t *testing.T) {
		box := structures.NewEmptyBox()
		boxItems := []structures.Item{
			{
				Weight:  10,
				Utility: 10,
			},
			{
				Weight:  5,
				Utility: 15,
			},
			{
				Weight:  20,
				Utility: 30,
			},
			{
				Weight:  15,
				Utility: 40,
			},
		}

		box.Items = boxItems

		backpack := structures.NewEmptyBackpack(20)
		backpackItems := []structures.Item{
			{
				Weight:  10,
				Utility: 1,
			},
			{
				Weight:  10,
				Utility: 5,
			},
		}
		backpack.Items = backpackItems

		newBackpack := Process(
			functions.ValidationFunction,
			functions.RandomSolution,
			10,
			100,
			100,
			backpack,
			box,
		)

		fmt.Println(newBackpack)
	})
}
