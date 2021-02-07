package main

import "testing"

func TestOla(t *testing.T) {
	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Cris", "")
		esperado := "Olá, Cris"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em espanhol", func(t *testing.T) {
		resultado := Ola("Eloide", "espanhol")
		esperado := "Hola, Eloide"

		verificarMensagemCorreta(t, resultado, esperado)
	})

}
