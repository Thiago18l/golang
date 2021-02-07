package loop

// Loop will return a sequence of same caracteres
func Loop(caractere string, quantity int) (word string) {
	for i := 0; i < quantity; i++ {
		word += caractere
	}
	return word
}
