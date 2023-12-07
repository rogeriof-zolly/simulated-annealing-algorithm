package functions

import (
	"anneling/src/structures"
	"math/rand"
)

// Solução aleatória gerada a partir da mochila e da caixa
func RandomSolution(bp structures.Backpack, box *structures.Box) structures.Backpack {
	var (
		boxPosition      int
		backpackPosition int
	)
selectItems:
	// Pega uma posição aleatória da caixa
	if len(box.Items) > 1 {
		boxPosition = rand.Intn(len(box.Items) - 1)
	} else {
		boxPosition = 0
	}

	// Pega uma posição aleatória da mochila
	if len(bp.GetItems()) > 1 {
		backpackPosition = rand.Intn(len(bp.GetItems()) - 1)
	} else {
		backpackPosition = 0
	}

	// Troca o item da mochila pelo item na caixa
	bp.Items[backpackPosition], box.Items[boxPosition] = box.Items[boxPosition], bp.Items[backpackPosition]

	// Se o item trocado faz a mochila passar do peso
	// Seleciona um novo item, voltando para o label de cima
	if bp.CalculateWeight() > bp.MaxWeight {
		goto selectItems
	}

	// Retorna a nova mochila
	return bp
}
