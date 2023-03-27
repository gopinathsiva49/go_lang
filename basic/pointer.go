package main

import "fmt"

func main() {
	var i int

	i = 1

	fmt.Println("i", i)

	var s *int

	s = &i

	fmt.Println("s:", s)
	// s = 2 // error
	*s = 2
	fmt.Println("s:", s)
	fmt.Println("i", i)

	//pointer value assign only pointer

}
