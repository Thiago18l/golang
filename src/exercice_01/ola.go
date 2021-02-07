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
	return prefixoSaudacao(lang) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo := prefixoOlafrances
		return prefixo
	case espanhol:
		prefixo := prefixoOlaEspanhol
		return prefixo
	default:
		prefixo := prefixoOlaPortugues
		return prefixo
	}
}

func main() {
	fmt.Println(Ola("", ""))
}
