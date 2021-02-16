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

	compareDefinition := func(t *testing.T, d Dictionary, word, definition string) {
		result, err := d.Busca(word)
		t.Helper()

		if err != nil {
			t.Fatal("Não foi possível encontrar a palavra adicionada", err)
		}
		verifyResult(t, definition, result)
	}
	dictionary := Dictionary{}
	word := "teste"
	definition := "isso é apenas um teste"

	dictionary.Adiciona(word, definition)
	compareDefinition(t, dictionary, word, definition)
}

func verifyResult(t *testing.T, waited, result string) {
	t.Helper()
	if waited != result {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}
