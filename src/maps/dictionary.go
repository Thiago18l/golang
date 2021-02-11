package maps

// Busca will return a string
func Busca(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func main() {
	mapWord := map[string]string{"teste": "isso é um teste"}
	Busca(mapWord, "teste")
}
