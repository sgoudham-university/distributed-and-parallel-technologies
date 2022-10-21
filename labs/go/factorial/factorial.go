package main

import (
	"fmt"
)

func factorial(n int64) int64 {
	if n == 0 {
		return 1;
	}
	return n * factorial(n -1)
}

func main() {
	fmt.Println("Factorial 16", factorial(16))
}
