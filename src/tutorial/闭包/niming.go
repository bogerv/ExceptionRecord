package main

func main() {
	f := func(x, y int) int {
		return x + y
	}
	f(1, 1)

	ch := make(chan int)
	func(ch chan int) {
		ch <- 9
	}(ch)
}
