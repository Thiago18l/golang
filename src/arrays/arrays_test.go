package arrays

import (
	"reflect"
	"testing"
)

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

func TestSomaTudo(t *testing.T) {
	verifySum := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
		}
	}

	t.Run("Sum all itens inside the slice", func(t *testing.T) {
		resultado := SomaTudo([]int{1, 2, 3}, []int{1, 1})
		esperado := []int{6, 2}
		verifySum(t, resultado, esperado)
	})

	t.Run("Sum empty slices", func(t *testing.T) {
		resultado := SomaTodoResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verifySum(t, resultado, esperado)
	})

	t.Run("Sum everything else", func(t *testing.T) {
		resultado := SomaTodoResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verifySum(t, resultado, esperado)
	})
}
