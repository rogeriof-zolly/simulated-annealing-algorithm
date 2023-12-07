package structures

import (
	"math/rand"
)

// Struct (objeto) Item
type Item struct {
	Weight  int
	Utility int
}

// Struct (objeto) Mochila
type Backpack struct {
	// Array de Items da mochila
	Items []Item
	// Capacidade máxima
	MaxWeight int
}

// Cria uma nova mochila vazia
func NewEmptyBackpack(maxWeight int) Backpack {
	return Backpack{[]Item{}, maxWeight}
}

// Enche a mochila com itens aleatórios da caixa recebida
func (backpack *Backpack) FillBackpack(box *Box) {
	var boxPosition int
	weight := 0

	// Enquanto o somador de peso não ultrapassar o peso máximo
	for weight < backpack.MaxWeight {

		// Pega uma posição aleatória da caixa
		if len(box.Items) > 1 {
			boxPosition = rand.Intn(len(box.Items) - 1)
		} else {
			boxPosition = 0
		}

		// Se o peso do item da caixa na posição + o somador de peso
		// passar a capacidade da mochila, sai do loop
		if box.Items[boxPosition].Weight+weight > backpack.MaxWeight {
			break
		}

		// Adiciona no acumulador de peso
		weight += box.Items[boxPosition].Weight

		// Adiciona o item na mochila
		backpack.Items = append(backpack.Items, box.Items[boxPosition])

		// Remove o item da caixa
		box.RemoveItem(boxPosition)
	}

}

// Adiciona item na mochila
func (backpack *Backpack) AddItem(w int, u int) {
	backpack.Items = append(backpack.Items, Item{w, u})
}

// Retorna um slice de itens da mochila
func (backpack *Backpack) GetItems() []Item {
	return backpack.Items
}

// Calcula o peso atual da mochila
func (backpack Backpack) CalculateWeight() int {
	sum := 0
	for _, item := range backpack.Items {
		sum += item.Weight
	}

	return sum
}
