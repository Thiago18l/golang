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
	verifyResult := func(t *testing.T, f Forma, waited float64) {
		t.Helper()
		result := f.Area()
		if result != waited {
			t.Errorf("result '%.2f', waited '%.2f'", result, waited)
		}
	}

	t.Run("Rentagles", func(t *testing.T) {
		retangle := Retangulo{12.0, 6.0}
		waited := 72.0

		verifyResult(t, retangle, waited)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		waited := 314.1592653589793

		verifyResult(t, circle, waited)
	})
}

func TestAreaTDT(t *testing.T) {
	testesArea := []struct {
		f      Forma
		waited float64
	}{
		{Retangulo{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range testesArea {
		result := tt.f.Area()
		if result != tt.waited {
			t.Errorf("result '%.2f', waited '%.2f'", result, tt.waited)
		}
	}
}
