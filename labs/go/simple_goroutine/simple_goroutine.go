package main

import "fmt"

func factorialGo(n int64, c chan int64) {
	c <- n
}

func main() {
	ch := make(chan int64)

	go factorialGo(20, ch)

	fmt.Println(<-ch)
}
