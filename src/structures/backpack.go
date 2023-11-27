package structures

type item struct {
	weight  int
	utility int
}

type Backpack struct {
	items []item
}

func NewEmptyBackpack() Backpack {
	return Backpack{[]item{}}
}

func (backpack *Backpack) addItem(w int, u int) {
	backpack.items = append(backpack.items, item{w, u})
}

func (backpack *Backpack) getItems() []item {
	return backpack.items
}
