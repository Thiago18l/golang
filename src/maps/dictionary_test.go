package maps

import "testing"

func TestBusca(t *testing.T) {
	verifyOutput := func(t *testing.T, result, waited, data string) {
		t.Helper()

		if result != waited {
			t.Errorf("result '%s', waited '%s', data '%s'", result, waited, data)
		}
	}

	dictionary := Dictionary{"teste": "isso é apenas um teste"}
	result := dictionary.Busca("teste")
	waited := "isso é apenas um teste"

	verifyOutput(t, result, waited, "teste")

}
