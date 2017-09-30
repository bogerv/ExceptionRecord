package main

import (
	"fmt"
)

func main() {
	// var val interface{} = int64(58)
	//// int64
	// fmt.Println(reflect.TypeOf(val))
	// val = 50
	//// int
	// fmt.Println(reflect.TypeOf(val))

	var val interface{} = nil
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}

	var val2 interface{} = (*interface{})(nil)
	// val2 = (*int)(nil)
	if val2 == nil {
		fmt.Println("val2 is nil")
	} else {
		fmt.Println("val2 is not nil")
	}
}
