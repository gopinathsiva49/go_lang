package main

import "fmt"

func main() {
	// In simple words, defer will move the execution of the statement to the very end inside a function.
	fmt.Println("hello")
	fmt.Println("hi")
	fmt.Println("welcom")
	fmt.Println("thanks")
	defer fmt.Println("worlds") // first enter into stack: run 3
	defer fmt.Println("world")  // second enter into stack: run 2
	defer stakingdefer()        // third enter into stack : run 1
}

// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

func stakingdefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
