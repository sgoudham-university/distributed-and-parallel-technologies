package main

import (
	"fmt"
	"sync"
	"time"
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

func perfectInterval(l, u int64, c chan int64, wg *sync.WaitGroup) {
	for i := l; i < u; i++ {
		if perfect(i) {
			c <- i
		}
	}

	wg.Done()
}

func main() {
	// Record start time
	start := time.Now()

	// Compute
	var wg sync.WaitGroup
	ch := make(chan int64, 10)

	for i := int64(0); i < 10000; i += 1000 {
		if i == 0 {
			i += 1
		}
		wg.Add(1)
		go perfectInterval(i, int64(i + 1000), ch, &wg)
	}

	wg.Wait()
	close(ch)

	for i := range ch {
		fmt.Println(i, "is perfect")
	}

	// Record the elapsed time
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed time", elapsed)
}
