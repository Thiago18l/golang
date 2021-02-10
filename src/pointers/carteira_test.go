package pointers

import "testing"

func TestCarteira(t *testing.T) {
	carteira := Carteira{}
	carteira.Depositar(10)

	result := carteira.Saldo()
	waited := 10

	if result != waited {
		t.Errorf("result '%d', waited '%d'", result, waited)
	}
}
