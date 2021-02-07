package loop

// Loop will return a sequence of same caracteres
func Loop(caractere string) (word string) {
	for i := 0; i < 5; i++ {
		word = word + caractere
	}
	return word
}
