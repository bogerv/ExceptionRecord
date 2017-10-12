package main

import "fmt"

func main(){
	var ptr *int
	num :=100
	ptr=&num
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr=200
	fmt.Println(*ptr)
	fmt.Println(num)
}