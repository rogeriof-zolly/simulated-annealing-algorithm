package structures

import "math/rand"

type item struct {
	weight  int
	utility int
}

type Backpack struct {
	items     []item
	maxWeight int
}

func NewEmptyBackpack() Backpack {
	return Backpack{[]item{}, 20}
}

func (backpack *Backpack) addItem(w int, u int) {
	backpack.items = append(backpack.items, item{w, u})
}

func (backpack *Backpack) getItems() []item {
	return backpack.items
}

func (backpack *Backpack) RandomSolution(box *Box) {
	var (
		boxPosition      int
		backpackPosition int
	)
	if len(box.items) > 1 {
		boxPosition = rand.Intn(len(box.items) - 1)
	} else {
		boxPosition = 0
	}

	if len(backpack.items) > 1 {
		backpackPosition = rand.Intn(len(backpack.items) - 1)
	} else {
		backpackPosition = 0
	}

	backpack.items[backpackPosition], box.items[boxPosition] = box.items[boxPosition], backpack.items[backpackPosition]
}

func (backpack Backpack) ValidationFunction() int {
	sum := 0

	for _, item := range backpack.items {
		sum += item.utility
	}

	return sum
}
