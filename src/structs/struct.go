package structs

// Retangulo sctruct
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Perimetro will return a float number
func Perimetro(x, y float64) float64 {
	return 2 * (x + y)
}

// Area will return a float number
func Area(base, altura float64) float64 {
	return base * altura
}
