package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// a (i)
// ch <- 3

// a (ii)
// fmt.Println(<-ch)

// a (iii)
// ch := make(chan int, 1)

// b (i) (ii)
// ch := make(chan int, 3)
// ch <- 1
// ch <- 2
// ch <- 3
// fmt.Println(<-ch)
// fmt.Println(<-ch)
// fmt.Println(<-ch)
