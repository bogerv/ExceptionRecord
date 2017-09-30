package main

import "fmt"

func change(num *int) {
    fmt.Println(num)    //0xc42000e1f8
    fmt.Println(*num)   //100
    *num = 1000
    fmt.Println(num)    //0xc42000e1f8
    fmt.Println(*num)   //1000
}

func main() {
    num := 100
    fmt.Println(&num)    //0xc42000e1f8
    change(&num)
    fmt.Println(&num)    //0xc42000e1f8
    fmt.Println(num)     //1000
}