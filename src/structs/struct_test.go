package structs

import "testing"

func TestPerimetro(t *testing.T) {
	result := Perimetro(10.0, 10.0)
	waited := 40.0

	if result != waited {
		t.Errorf("result '%.2f', waited '%.2f'", result, waited)
	}
}

func TestArea(t *testing.T) {
	result := Area(12.0, 6.0)
	waited := 72.0

	if result != waited {
		t.Errorf("result '%.2f', waited '%.2f'", result, waited)
	}
}
