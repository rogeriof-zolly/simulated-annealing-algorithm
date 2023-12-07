package main

import (
	"anneling/src/functions"
	"anneling/src/procedure"
	"anneling/src/testcases"
	"fmt"
)

func testeGrande() (int, int) {
	fmt.Println("Teste grande")
	bo, ba := testcases.ReadFromFile("./src/testcases/files/large_scale/knapPI_3_10000_1000_1")
	resultadoEsperado := testcases.ReadExpectedResultFromFile("./src/testcases/files/large_scale-optimum/knapPI_3_10000_1000_1")

	ba.FillBackpack(&bo)
	value := procedure.Process(
		functions.ValidationFunction,
		functions.RandomSolution,
		0.2,
		1000,
		500,
		ba,
		bo,
	)
	return functions.ValidationFunction(value), resultadoEsperado
}

func testePequeno() (int, int) {
	fmt.Println("Teste pequeno")
	bo, ba := testcases.ReadFromFile("./src/testcases/files/low-dimensional/f10_l-d_kp_20_879")
	resultadoEsperado := testcases.ReadExpectedResultFromFile("./src/testcases/files/low-dimensional-optimum/f10_l-d_kp_20_879")

	ba.FillBackpack(&bo)
	value := procedure.Process(
		functions.ValidationFunction,
		functions.RandomSolution,
		0.2,
		1000,
		100,
		ba,
		bo,
	)
	return functions.ValidationFunction(value), resultadoEsperado
}

func main() {

	resultadoTG, resultadoEsperadoTG := testeGrande()
	fmt.Println("Resultado do teste grande: ", resultadoTG)
	fmt.Println("Resultado esperado do teste grande: ", resultadoEsperadoTG)

	resultadoTP, resultadoEsperadoTP := testePequeno()
	fmt.Println("Resultado do teste pequeno: ", resultadoTP)
	fmt.Println("Resultado esperado do teste pequeno: ", resultadoEsperadoTP)

	fmt.Println("Pressione enter para encerrar a execução")
	fmt.Scanln()
}
