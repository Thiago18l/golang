package maps

// Dictionary type
type Dictionary map[string]string

// Busca will return a string
func (d Dictionary) Busca(word string) string {
	return d[word]
}

func main() {
	dict := Dictionary{"teste": "isso aqui é um teste"}
	dict.Busca("teste")
}
