package maps

import "testing"

func TestBusca(t *testing.T) {
	dictionary := Dictionary{"teste": "isso é apenas um teste"}

	verifyOutput := func(t *testing.T, result, waited, data string) {
		t.Helper()

		if result != waited {
			t.Errorf("result '%s', waited '%s', data '%s'", result, waited, data)
		}
	}
	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Busca("teste")
		waited := "isso é apenas um teste"

		verifyOutput(t, result, waited, "teste")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, result := dictionary.Busca("Unknown")
		compareError(t, result, ErrNotFound)
	})
}

func compareError(t *testing.T, result, waited error) {
	t.Helper()
	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}

func TestAdiciona(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Adiciona("teste", "isso é apenas um teste")
	waited := "isso é apenas um teste"

	result, err := dictionary.Busca("teste")
	if err != nil {
		t.Fatal("Não foi possível encontrar a palavra adicionada", err)
	}
	if waited != result {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}
