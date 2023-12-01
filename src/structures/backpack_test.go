package structures

import "testing"

func TestBackpack(t *testing.T) {
	t.Run("Should add item to backpack", func(t *testing.T) {
		backpack := NewEmptyBackpack()

		backpack.AddItem(10, 40)

		expected := Item{Weight: 10, Utility: 40}

		if backpack.GetItems()[0] != expected {
			t.Error("Test failed")
		}
	})
}
