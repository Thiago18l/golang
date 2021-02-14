package maps

import "errors"

// ErrNotFound details
var ErrNotFound = errors.New("não foi possível encontrar a palavra que você procura")

// Dictionary type
type Dictionary map[string]string

// Busca will return a string
func (d Dictionary) Busca(word string) (string, error) {
	definicao, existe := d[word]
	if !existe {
		return "", ErrNotFound
	}
	return definicao, nil
}

func main() {
	dict := Dictionary{"teste": "isso aqui é um teste"}
	dict.Busca("teste")
}
