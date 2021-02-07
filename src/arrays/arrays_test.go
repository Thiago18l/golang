package arrays

import "testing"

func TestSoma(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	resultado := Soma(numbers)
	esperado := 15
	if esperado != resultado {
		t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
	}
}
