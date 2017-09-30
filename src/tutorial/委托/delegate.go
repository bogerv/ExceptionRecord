package main

import (
	"fmt"
	"reflect"
)

type addFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op addFunc, a int, b int) int {
	return op(a, b)
}

func main() {
	c := add
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
	sum := operator(c, 1, 2)
	fmt.Println(sum)
}
