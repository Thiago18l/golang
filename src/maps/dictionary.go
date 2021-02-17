package maps

import "errors"

// ErrNotFound details
var (
	ErrNotFound     = errors.New("não foi possível encontrar a palavra que você procura")
	ErrExistingWord = errors.New("não é possível adicionar a palavra pois ela já existe")
)

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

// Adiciona will add a new word to the map
func (d Dictionary) Adiciona(word, definition string) error {
	_, err := d.Busca(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrExistingWord
	default:
		return err
	}
	return nil
}

func main() {
	dict := Dictionary{"teste": "isso aqui é um teste"}
	dict.Busca("teste")
}
