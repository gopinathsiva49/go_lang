package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func add(a int, b int) int {
	return a + b
}

func multiple(a, b int) int {
	return a * b
}

func multiresult(a, b string) (string, string) {
	return a, b
}

func variablevalue(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}

func withstruct(data Student) (string, int, Student) {
	return data.Name, data.Age, data
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(multiple(2, 2))

	// multi result with variable
	a, b := multiresult("hi", "welcome")
	fmt.Println(a, b)

	// Note that: := is a declaration, whereas = is an assignment
	// ref - https://stackoverflow.com/questions/36512919/what-is-the-difference-between-and-in-go
	fmt.Println(variablevalue(1, 2))

	x := Student{"gopinath", 1}
	fmt.Println(withstruct(x))
}
