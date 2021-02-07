package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

// Ola retorna um string
func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola(""))
}
