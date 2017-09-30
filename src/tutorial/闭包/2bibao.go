package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f1 := makeSuffix(".png")
	fmt.Println(f1("name1"))
	fmt.Println(f1("name2"))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("name1"))
	fmt.Println(f2("name2"))
}
