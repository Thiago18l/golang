package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Cumprimenta func
func Cumprimenta(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Olá, %s", name)
}

// HandleMyMeeting output for web
func HandleMyMeeting(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "Mundo")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandleMyMeeting))
	Cumprimenta(os.Stdout, "Thiago")
	if err != nil {
		fmt.Println(err)
	}
}
