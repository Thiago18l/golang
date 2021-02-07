package loop

import "testing"

func TestLoop(t *testing.T) {
	repeticoes := Loop("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("resultado '%s', esperado '%s'", repeticoes, esperado)
	}
}
