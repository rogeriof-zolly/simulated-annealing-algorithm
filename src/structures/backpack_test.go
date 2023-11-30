package structures

import "testing"

func TestBackpack(t *testing.T) {
	t.Run("Should add item to backpack", func(t *testing.T) {
		backpack := NewEmptyBackpack()

		backpack.addItem(10, 40)

		expected := item{weight: 10, utility: 40}

		if backpack.getItems()[0] != expected {
			t.Error("Test failed")
		}
	})

	t.Run("Should switch positions between box and backpack", func(t *testing.T) {
		bp := NewEmptyBackpack()
		b := NewEmptyBox()

		bp.addItem(10, 40)
		b.addItem(30, 4)

		backpackItemBefore := bp.items[0]
		boxItemBefore := b.items[0]

		bp.RandomSolution(&b)

		if bp.items[0] != boxItemBefore {
			t.Error("Failed to change backpack item")
		}

		if b.items[0] != backpackItemBefore {
			t.Error("Failed to change backpack item")
		}
	})
}
