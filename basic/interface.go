package main

import "fmt"

type Docter struct {
	Id   int
	Name string
}

func (d *Docter) first() {
	d.Name = "gopinath"
}

func main() {

	d := &Docter{}
	d.first()
	// d.Name = "gopinath"
	fmt.Println(d)
}

// interface in class intance function
