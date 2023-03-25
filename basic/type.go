package main

import (
	"fmt"
	"math/cmplx"
)

// bool
// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

var (
	t  bool       = true
	mi uint64     = 1<<64 - 1
	z  complex128 = cmplx.Sqrt(-5 + 12i)
	tt bool
)

func main() {
	fmt.Println(t)
	fmt.Println(mi)
	fmt.Println(z)
	fmt.Println(tt)

	// Variables declared without an explicit initial value are given their zero value.

	// The zero value is:

	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.

	var i int
	var f float64
	var b bool
	var s string

	fmt.Println(i, f, b, s)

	i = 5
	fmt.Println(i)

	// type conversion
	var y int
	fmt.Println(y)
	var z uint = uint(y)
	fmt.Println(z)

	// type interferce

	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	// When the right hand side of the declaration is typed, the new variable is of that same type:
	var c int
	d := c // d is an int
	fmt.Println(d)

	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
	ii := 42           // int
	ff := 3.142        // float64
	gg := 0.867 + 0.5i // complex128
	fmt.Println(ii, ff, gg)
}
