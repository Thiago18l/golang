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

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "teste"
		definition := "isso é apenas um teste"

		err := dictionary.Adiciona(word, definition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "teste"
		definition := "isso é apenas um teste"
		dictionary := Dictionary{word: definition}
		err := dictionary.Adiciona(word, "teste novo")
		compareError(t, err, ErrExistingWord)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "teste"
	definition := "isso é apenas um teste"
	dictionary := Dictionary{word: definition}
	newDefinition := "nova definição"

	dictionary.Atualiza(word, newDefinition)
	compareDefinition(t, dictionary, word, newDefinition)
}

func verifyResult(t *testing.T, waited, result string) {
	t.Helper()
	if waited != result {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}

func compareDefinition(t *testing.T, d Dictionary, word, definition string) {
	result, err := d.Busca(word)
	t.Helper()

	if err != nil {
		t.Fatal("Não foi possível encontrar a palavra adicionada", err)
	}
	verifyResult(t, definition, result)
}
