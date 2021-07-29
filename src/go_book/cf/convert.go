// TO RUN THIS FILE IS NECESSARY TO MAKE PUBLIC THE FUNCTIONS OF THE PACKAGE FTOC AND CHANGE THE NAME OF THE PACKAGE TO
// ftoc
package main

import (
	"fmt"
	"os"
	"src/maps/src/go_book/ftoc"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := ftoc.Fahrenheit(t)
		c := ftoc.Celsius(t)
		fmt.Printf("%g = %g, %g = %g\n", f, ftoc.FToC(f), c, ftoc.CtoF(c))
	}
}
