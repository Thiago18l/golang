package structs

import "math"

// Retangulo struct
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Circle struct
type Circle struct {
	Raio float64
}

// Perimetro will return a float64
func Perimetro(r Retangulo) float64 {
	return 2 * (r.Largura + r.Altura)
}

// Area method for Rentagulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Area will return a float number
func (c Circle) Area() float64 {
	return math.Pi * (math.Pow(c.Raio, 2))
}
