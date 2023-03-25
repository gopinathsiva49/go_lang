package main

import "fmt"

// Go has only one looping construct, the for loop.
// Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
	forcontinueorwhile()
	forever()
	rangeloop()
}

func forcontinueorwhile() {
	fmt.Println("forcontinueorwhile example")
	sum := 1
	for sum < 10000 {
		sum += 1
		fmt.Println(sum)
	}
}

func forever() {
	fmt.Println("forever example")
	// for {
	// 	fmt.Println("test")
	// }
}

func rangeloop() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Println(i, v)
	}
	for _, value := range pow {
		fmt.Println(value)
	}
}
