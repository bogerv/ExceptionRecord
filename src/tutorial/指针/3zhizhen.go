package main

import (
	"fmt"
)

type Rect struct {
	width, height float64
}

func (rect *Rect) area() float64 {
	rect.width = 200
	rect.height = 300
	return rect.width * rect.height
}

func main() {
	var rect = new(Rect)
	rect.width = 100
	rect.height = 200
	fmt.Println(*rect)
	fmt.Println("Width:", rect.width, "Length:", rect.height, "Area:", rect.area())
	fmt.Println(*rect)
}
