package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func main() {
	fmt.Println(MyError{time.Now(), "it didn't work"})
}
