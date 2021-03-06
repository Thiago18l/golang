package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("resultado '%d', esperado '%d'", soma, esperado)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 6)
	fmt.Println(soma)
	// Output: 7
}
