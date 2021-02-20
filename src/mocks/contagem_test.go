package mocks

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}

	Contagem(buffer)

	result := buffer.String()
	waited := "3"
	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}
