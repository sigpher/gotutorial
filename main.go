package main

import (
	"fmt"
	"strings"
)

func main() {
	op := adder()
	fmt.Println(op(10))
	fmt.Println(op(10))
	fmt.Println(op(10))
	fmt.Println(op(10))
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
