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

	i := 1
	increment(&i)              // side effect: i is now 2
	fmt.Println(increment(&i)) // 3 and i is 3
}

func f() *int {
	v := 1
	return &v
}

func increment(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
