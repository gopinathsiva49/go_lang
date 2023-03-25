package main

import "fmt"

// in go variable always have default value

var c, python, java bool
var s string

func main() {
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(s)

	var a, b int = 1, 2 // with type
	fmt.Println(a, b)

	var d, e, f = true, 1, "test" // without type
	fmt.Println(d, e, f)
	shortv()
	constants()
}

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
// h := 1 // error

func shortv() {
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	g := 1
	fmt.Println(g)
	// g := 2 // error
	g = 25
	fmt.Println(g)

	f := 1
	fmt.Println(f)
}

func constants() {
	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.

	const Pi = 3.14
	fmt.Println(Pi)
	// Pi = 2.13 // unable to change

}
