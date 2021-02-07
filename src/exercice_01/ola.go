package main

import "fmt"

// Ola retorna um string
func Ola(nome string) string {
	return "Olá, " + nome
}

func main() {
	fmt.Println(Ola("nome"))
}
