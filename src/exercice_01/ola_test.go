package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Cris")
	esperado := "Olá, Cris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
