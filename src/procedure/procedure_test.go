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
			{Weight: 10, Utility: 4},
			{Weight: 47, Utility: 60},
			{Weight: 5, Utility: 32},
			{Weight: 4, Utility: 23},
			{Weight: 50, Utility: 72},
			{Weight: 8, Utility: 80},
			{Weight: 87, Utility: 46},
			{Weight: 85, Utility: 65},
			{Weight: 55, Utility: 95},
		}

		box.Items = boxItems

		backpack := structures.NewEmptyBackpack(269)
		backpack.FillBackpack(&box)

		fmt.Println(backpack.Items)

		value := Process(
			functions.ValidationFunction,
			functions.RandomSolution,
			0.3,
			200,
			100,
			backpack,
			box,
		)

		fmt.Println(functions.ValidationFunction(value))
	})
}
