package main

import "fmt"

func main() {
	a1 := []int{3, 2, 4}
	var r1, r2 int
	var res int
	for key, value := range a1 {
		numsRight := a1[key+1 : len(a1)]
		fmt.Println(numsRight)
		fmt.Println("------")
		res = key + 1
		for keyRight, valueRight := range numsRight {
			if (value + valueRight) == 6 {
				r1 = key
				r2 = keyRight + res
			}
		}
	}
	fmt.Println(r1)
	fmt.Println(r2)
}
