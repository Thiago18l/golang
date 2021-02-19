package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

// Cumprimenta func
func Cumprimenta(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Olá, %s", name)
}

func main() {
	Cumprimenta(os.Stdout, "Thiago")
}
