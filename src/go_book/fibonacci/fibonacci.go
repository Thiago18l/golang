package main

import "fmt"

func main() {
	result := fibonacci(10)
	fmt.Println(result)
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
