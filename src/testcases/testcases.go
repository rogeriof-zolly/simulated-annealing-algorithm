package testcases

import (
	"anneling/src/structures"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Lê um arquivo e retorna uma caixa de itens e uma mochila vazia com o tamanho definido no arquivo
func ReadFromFile(filename string) (structures.Box, structures.Backpack) {
	// Importação do arquivo
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Comando defer atrasa a execução da função ao lado
	// Neste caso, no último momento possível, em return ou erro
	// o arquivo é fechado
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Controle do número da linha
	lineNumber := 0

	// Definição de variáveis
	var (
		bo      structures.Box
		ba      structures.Backpack
		column1 int
		column2 int
	)

	// Iteração sobre as linhas
	for scanner.Scan() {
		lineNumber += 1

		// Divisão dos dados da linha
		splittedLine := strings.Split(scanner.Text(), " ")

		column1, err = strconv.Atoi(splittedLine[0])
		if err != nil {
			log.Fatal(err)
		}

		column2, err = strconv.Atoi(splittedLine[1])
		if err != nil {
			log.Fatal(err)
		}

		// Se for a primeira linha, define a mochila com o peso do arquivo
		if lineNumber == 1 {
			ba = structures.NewEmptyBackpack(column2)
			continue
		}

		// Adiciona na caixa de itens disponíveis o item com o peso e a utilidade
		bo.Items = append(bo.Items, structures.Item{Weight: column1, Utility: column2})

	}

	// Retorna a caixa e a mochila
	return bo, ba
}

func ReadExpectedResultFromFile(filename string) int {
	// Importação do arquivo
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expected int

	for scanner.Scan() {
		// Converte para inteiro o valor da linha
		expected, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	// Retorna o valor esperado definido no arquivo
	return expected
}
