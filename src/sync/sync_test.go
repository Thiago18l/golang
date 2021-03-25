package main

import "testing"

func TestCount(t *testing.T) {
	t.Run("Increment the count 3 times result in third value", func(t *testing.T) {
		contador := Contador{}
		contador.Increment()
		contador.Increment()
		contador.Increment()

		if contador.Value() != 3 {
			t.Errorf("result %d, waited %d", contador.Value(), 3)
		}
	})
}
