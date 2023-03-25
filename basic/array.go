package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	sa := [5]Student{}
	fmt.Println(sa)
	slicearr()
}

func slicearr() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	fmt.Println("----------------------")
	fmt.Println(q[1:4])
	fmt.Println(q[:2])
	fmt.Println(q[1:])
	fmt.Println("----------------------")

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	y := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(y)

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

	// The make function allocates a zeroed array and returns a slice that refers to that array:

	// a := make([]int, 5)  // len(a)=5

	z := make([]int, 5)
	fmt.Println(z)
	fmt.Println(len(z))

}
