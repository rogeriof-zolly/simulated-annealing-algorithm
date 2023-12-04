package structures

import "math/rand"

type Item struct {
	Weight  int
	Utility int
}

type Backpack struct {
	Items     []Item
	maxWeight int
}

func NewEmptyBackpack(maxWeight int) Backpack {
	return Backpack{[]Item{}, maxWeight}
}

func (backpack *Backpack) AddItem(w int, u int) {
	backpack.Items = append(backpack.Items, Item{w, u})
}

func (backpack *Backpack) GetItems() []Item {
	return backpack.Items
}

func (backpack *Backpack) RandomSolution(box *Box) {
	var (
		boxPosition      int
		backpackPosition int
	)
	if len(box.Items) > 1 {
		boxPosition = rand.Intn(len(box.Items) - 1)
	} else {
		boxPosition = 0
	}

	if len(backpack.Items) > 1 {
		backpackPosition = rand.Intn(len(backpack.Items) - 1)
	} else {
		backpackPosition = 0
	}

	backpack.Items[backpackPosition], box.Items[boxPosition] = box.Items[boxPosition], backpack.Items[backpackPosition]
}

func (backpack Backpack) ValidationFunction() int {
	sum := 0

	for _, item := range backpack.Items {
		sum += item.Utility
	}

	return sum
}
