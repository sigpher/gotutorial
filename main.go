package main

import "fmt"

func main() {
	var x Mover
	var d1 = Dog{}
	x = d1
	x.Move()
	var d2 = &Dog{}
	x = d2
	x.Move()

	var c1 = &Cat{}
	x = c1
	x.Move()

}

type Mover interface {
	Move()
}

type Dog struct {
}

func (d Dog) Move() {
	fmt.Println("Dog running")
}

type Cat struct{}

func (c *Cat) Move() {
	fmt.Println("Cat running")
}
