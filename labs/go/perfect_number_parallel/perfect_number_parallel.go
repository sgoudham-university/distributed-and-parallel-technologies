package main

import (
	"fmt"
)

func perfect(n int64) bool {
	var running_count int64

	// special case:
	// loop will return true so handle initial case of n = 1 outside
	if n == 1 {
		return false
	}

	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			running_count += i
		}
	}

	return running_count == n
}

func perfectInterval(l, u int64, c chan int64) {
	for i := l; i < u; i++ {
		if perfect(i) {
			c <- i
		}
	}

	close(c)
}

func main() {
	ch := make(chan int64)

	for i := int64(0); i < 10000; i += 1000 {
		if i == 0 {
			i += 1
		}
		go perfectInterval(i, int64(i + 1000), ch)
	}

	// Currently, the main thread finishes before all the
	// goroutines have had a chance to finish

	// Unsure of how to ensure all goroutines are ran without
	// modifying the function signature of perfectInterval()
	for i := range ch {
		fmt.Println(i, "is perfect")
	}
}
