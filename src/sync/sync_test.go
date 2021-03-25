package main

import "testing"

func TestCount(t *testing.T) {

	verifyCount := func(t *testing.T, result Contador, waited int) {
		t.Helper()
		if result.Value() != waited {
			t.Errorf("result %d, waited %d", result.Value(), waited)
		}
	}
	t.Run("Increment the count 3 times result in third value", func(t *testing.T) {
		contador := Contador{}
		contador.Increment()
		contador.Increment()
		contador.Increment()

		verifyCount(t, contador, 3)
	})
}
