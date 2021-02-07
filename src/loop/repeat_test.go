package loop

import "testing"

func TestLoop(t *testing.T) {
	repeticoes := Loop("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("resultado '%s', esperado '%s'", repeticoes, esperado)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("a", 5)
	}
}
