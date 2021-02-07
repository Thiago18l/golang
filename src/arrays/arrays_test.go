package arrays

import "testing"

func TestSoma(t *testing.T) {

	result := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
		}
	}

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		resultado := Soma(numbers)
		esperado := 15
		result(t, resultado, esperado)
	})
}
