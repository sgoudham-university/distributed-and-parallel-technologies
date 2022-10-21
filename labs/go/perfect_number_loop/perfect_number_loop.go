package main

import (
	"fmt"
	"time"
)

// n1 = 28, LPD=1,2  281=28, 142=28
// n2 = 36, LPD=1,2  361=36, 182=36
// n3 = 45, LPD=1,5, 451=45, 95=45
//
// The number itself should not be included
// We know that the lowest possible proper divisors are 1 and 2
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

func perfectLoop(input int64) {
	for i := int64(1); i <= input; i++ {
		if perfect(i) {
			fmt.Println(i, "is perfect")
		}
	}
}

func perfectLoopInterval(l,u int64) {
	for i := l; i <= u; i++ {
		if perfect(i) {
			fmt.Println(i, "is perfect")
		}
	}
}

func main() {
	// Record start time
	start := time.Now()

	// Compute
	for i := int64(0); i < 10000; i += 1000 {
		if i == 0 {
			i += 1
		}
		perfectLoopInterval(i, int64(i + 1000))
	}

	// Record the elapsed time
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed time", elapsed)
}


/*
2.5: Performance Measurements

1000 - 826.41µs
2000 - 4.170284ms
4000 - 10.942887ms
8000 - 37.615378ms
16000 - 2.05558ms
*/
