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
}
