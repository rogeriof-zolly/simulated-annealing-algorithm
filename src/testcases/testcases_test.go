package testcases

import (
	"anneling/src/functions"
	"anneling/src/procedure"
	"fmt"
	"testing"
)

func TestReadingFiles(t *testing.T) {

	t.Run("Reading large file", func(t *testing.T) {
		bo, ba := ReadFromFile("/large_scale/knapPI_3_10000_1000_1")

		ba.FillBackpack(&bo)
		fmt.Println(bo)
		fmt.Println(ba)
		value := procedure.Process(
			functions.ValidationFunction,
			functions.RandomSolution,
			0.2,
			1000,
			100,
			ba,
			bo,
		)
		fmt.Println(functions.ValidationFunction(value))
	})
}
