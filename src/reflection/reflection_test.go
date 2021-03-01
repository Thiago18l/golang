package main

import "testing"

func TestPercorre(t *testing.T) {
	waited := "Thiago"
	var result []string
	x := struct {
		Name string
	}{waited}

	percorre(x, func(entrada string) {
		result = append(result, entrada)
	})
	if result[0] != waited {
		t.Errorf("número incorreto de chamadas de função: result '%s', waited '%s'", result[0], waited)
	}
}
