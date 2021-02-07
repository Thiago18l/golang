package main

import "fmt"

const espanhol = "espanhol"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

// Ola retorna um string
func Ola(nome string, lang string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if lang == espanhol {
		return prefixoOlaEspanhol + nome
	}

	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("", ""))
}
