package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	way()
	diff()
}

// the new function

func way() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func diff() {
	p := new(int)
	q := new(int)
	fmt.Println(p == q) // false
}
