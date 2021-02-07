package inteiros

import "fmt"

// Adiciona retorna a soma de dois números
func Adiciona(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(Adiciona(2, 2))
}
