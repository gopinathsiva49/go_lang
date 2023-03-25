package main

import "fmt"

type Student struct {
	Id      int
	Name    string
	Age     int
	Address string
}

func main() {
	fmt.Println(Student{1, "gopinath", 26, "kvb garden ran puram chennai - 28"})

	first := Student{1, "gopinath", 26, "kvb garden ran puram chennai - 28"}
	fmt.Println(first)
	first.Name = "gopinath sivakumar"
	fmt.Println(first)

	second := Student{Name: "gopi"}
	fmt.Println(second)
}
