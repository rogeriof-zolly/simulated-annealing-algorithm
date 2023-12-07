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

func (box *Box) RemoveItem(index int) {
	box.Items[index] = box.Items[len(box.Items)-1]
	box.Items = box.Items[:len(box.Items)-1]
}
