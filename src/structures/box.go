package structures

type Box struct {
	items []item
}

func NewEmptyBox() Box {
	return Box{[]item{}}
}

func (box *Box) addItem(w int, u int) {
	box.items = append(box.items, item{w, u})
}
