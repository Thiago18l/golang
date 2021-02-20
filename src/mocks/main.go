package mocks

import (
	"fmt"
	"io"
	"os"
)

// Contagem will print a output in screen
func Contagem(exit io.Writer) {
	fmt.Fprintf(exit, "3")
}

func main() {
	Contagem(os.Stdout)
}
