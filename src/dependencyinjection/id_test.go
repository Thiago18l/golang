package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Thiago")
	result := buffer.String()
	waited := "Olá, Thiago"

	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}
