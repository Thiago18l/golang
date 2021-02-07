package arrays

import "testing"

func TestSoma(t *testing.T) {

	result := func(t *testing.T, resultado int, esperado int, dados []int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dados '%v'", resultado, esperado, dados)
		}
	}

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		resultado := Soma(numbers)
		esperado := 15
		result(t, resultado, esperado, numbers)
	})

	t.Run("Collection of any size", func(t *testing.T) {
		array := []int{1, 2, 3}
		resultado := Soma(array)
		esperado := 6
		result(t, resultado, esperado, array)
	})
}
