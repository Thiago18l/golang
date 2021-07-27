package main

import "fmt"

var p = f()

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2 // equivalent to x = 2
	fmt.Println(x)

	var j, k int
	fmt.Println(&j == &j, &j == &k, &x == nil)
	fmt.Println(f() == f())
}

func f() *int {
	v := 1
	return &v
}
