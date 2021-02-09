package structs

// Retangulo sctruct
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Perimetro will return a float number
func Perimetro(r Retangulo) float64 {
	return 2 * (r.Altura + r.Largura)
}

// Area will return a float number
func Area(r Retangulo) float64 {
	return r.Largura * r.Altura
}
