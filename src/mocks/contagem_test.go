package mocks

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}
	Contagem(buffer, sleeperSpy)

	result := buffer.String()
	waited := `3
	2
	1
	Go!`
	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
	if sleeperSpy.Chamadas != 4 {
		t.Errorf("Não houve chamadas suficientes do sleeper, waited 4, result '%d'", sleeperSpy.Chamadas)
	}
}
