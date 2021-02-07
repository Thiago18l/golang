package main

import "fmt"

const frances = "frances"
const espanhol = "espanhol"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlafrances = "Bonjour, "

// Ola retorna um string
func Ola(nome string, lang string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if lang == espanhol {
		return prefixoOlaEspanhol + nome
	}
	if lang == frances {
		return prefixoOlafrances + nome
	}

	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("", ""))
}
