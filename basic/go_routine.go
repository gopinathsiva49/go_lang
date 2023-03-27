package main

import (
	"fmt"
	"time"
)

func ss() {
	fmt.Println("first")
}
func main() {
	go ss()
	fmt.Println("second_first")
	fmt.Println("third_first")
	time.Sleep(time.Second)
	fmt.Println("four")
}
