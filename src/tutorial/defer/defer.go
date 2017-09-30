package main

import "fmt"

func test1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func test2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	test1()
	test2()
}
