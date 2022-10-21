package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func perfect(n int64) bool {
	var running_count int64

	// special case:
	// loop will return true so handle initial case of n = 1 outside
	if n == 1 {
		return false
	}

	for i := int64(1); running_count != n; i++ {
		if running_count > n {
			return false
		} else if n%i == 0 {
			running_count += i
		}
	}

	return true
}

func main() {
	var input int64
	var err error

	if len(os.Args) < 2 {
		panic(fmt.Sprintf("Usage: must provide number as an argument"))
	}

	if input, err = strconv.ParseInt(os.Args[1], 10, 64); err != nil {
		panic(fmt.Sprintf("Can't parse the first argument"))
	}

	// Record start time
	start := time.Now()
	// Compute
	fmt.Println(input, "is", perfect(input))
	// Record the elapsed time
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed time", elapsed)
}
