package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(findthegreater(1, 2))
	fmt.Println(findthegreater(2, 1))
	fmt.Println(findthegreater(2, 2))
	fmt.Println(fiwithshortvariable())
	switchcase()
	switchevalution()
	switchwithnocond()
}

func findthegreater(a, b int) string {
	if a < b {
		return "b is greater"
	} else if a == b {
		return "a and b are equal"
	} else {
		return "a is greater"
	}
}

func fiwithshortvariable() int {
	if v := 2; v == 1 {
		return v
	} else if v == 2 {
		return v
	}
	return 3
}

func switchcase() {
	fmt.Println("switchcase")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchevalution() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchwithnocond() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
