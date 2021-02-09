package structs

import "testing"

func TestPerimetro(t *testing.T) {
	retangle := Retangulo{10.0, 10.0}
	result := Perimetro(retangle)
	waited := 40.0

	if result != waited {
		t.Errorf("result '%.2f', waited '%.2f'", result, waited)
	}
}

func TestArea(t *testing.T) {
	retangle := Retangulo{12.0, 6.0}
	result := Area(retangle)
	waited := 72.0

	if result != waited {
		t.Errorf("result '%.2f', waited '%.2f'", result, waited)
	}
}
