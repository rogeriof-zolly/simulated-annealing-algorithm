package structures

type Box struct {
	Items []Item
}

func NewEmptyBox() Box {
	return Box{[]Item{}}
}

func (box *Box) AddItem(w int, u int) {
	box.Items = append(box.Items, Item{w, u})
}
