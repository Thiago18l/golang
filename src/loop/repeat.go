package loop

const repeticoes = 5

// Loop will return a sequence of same caracteres
func Loop(caractere string) (word string) {
	for i := 0; i < repeticoes; i++ {
		word += caractere
	}
	return word
}
