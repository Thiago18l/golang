package maps

import "errors"

// ErrNotFound details
var (
	ErrNotFound       = errors.New("não foi possível encontrar a palavra que você procura")
	ErrExistingWord   = errors.New("não é possível adicionar a palavra pois ela já existe")
	ErrInexistingWord = errors.New("não foi possível atualizar a palavra pois ela não existe")
)

// ErrDictionary type
type ErrDictionary string

// Dictionary type
type Dictionary map[string]string

func (e ErrDictionary) Error() string {
	return string(e)
}

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

// Atualiza will update a word in map
func (d Dictionary) Atualiza(word, definition string) error {
	_, err := d.Busca(word)
	switch err {
	case ErrNotFound:
		return ErrInexistingWord
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete is a func to delete the word inside the map
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
func main() {
	dict := Dictionary{"teste": "isso aqui é um teste"}
	dict.Busca("teste")
}
