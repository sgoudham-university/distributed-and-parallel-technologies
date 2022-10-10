package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Go >1.18 Generics
type Pair[T, U any] struct {
	First  T
	Second U
}

func getArray() []int {
	length := len(os.Args)

	numbers := make([]int, length-1)
	for i := 1; i < length; i++ {
		j, err := strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}
		numbers[i-1] = j
	}

	return numbers
}

func average(arr []int, channel chan Pair[string, int]) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	channel <- Pair[string, int]{"Average:", sum / len(arr)}
}

func sum(arr []int, channel chan Pair[string, int]) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	channel <- Pair[string, int]{"Sum:", sum}
}

func product(arr []int, channel chan Pair[string, int]) {
	product := arr[0]
	for i := 1; i < len(arr); i++ {
		product *= arr[i]
	}
	channel <- Pair[string, int]{"Product:", product}
}

func max(arr []int, channel chan Pair[string, int]) {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	channel <- Pair[string, int]{"Max:", max}
}

func min(arr []int, channel chan Pair[string, int]) {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	channel <- Pair[string, int]{"Min:", min}
}

func first(arr []int, channel chan Pair[string, int]) {
	channel <- Pair[string, int]{"First Element:", arr[0]}
}

func last(arr []int, channel chan Pair[string, int]) {
	channel <- Pair[string, int]{"Last Element:", arr[len(arr)-1]}
}

func print(value Pair[string, int], wg *sync.WaitGroup) {
	fmt.Println(value.First, value.Second)
	wg.Done()
}

// I know what you're thinking... this is a pretty dumb program
//
// Guess what... you're right!
//
// The point of this code was to get used to passing around data within channels and calling goroutines
// And trying to synchronise them using **sync.WaitGroup**
func main() {
	arr := getArray()
	fmt.Println("Original Array:", arr)
	fmt.Println()

	var wg sync.WaitGroup
	channel := make(chan Pair[string, int])

	go average(arr, channel)
	go sum(arr, channel)
	go product(arr, channel)
	go min(arr, channel)
	go max(arr, channel)
	go first(arr, channel)
	go last(arr, channel)

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go print(<-channel, &wg)
	}

	wg.Wait()
	fmt.Println("\nAll Array Operations Printed!")
}
