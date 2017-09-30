package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(3)
	numberChan1 := make(chan int64, 3)
	numberChan2 := make(chan int64, 3)
	numberChan3 := make(chan int64, 3)

	go func() {
		for n := range numberChan1 {
			if n%2 == 0 {
				numberChan2 <- n
			} else {
				fmt.Printf("Filter %d. [filter 1]\n", n)
			}
		}
		close(numberChan2)
		waitGroup.Done()
	}()

	go func() {
		for n := range numberChan2 {
			if n%5 == 0 {
				numberChan3 <- n
			} else {
				fmt.Printf("Filter %d. [filter 1]\n", n)
			}
		}
		close(numberChan3)
		waitGroup.Done()
	}()

	go func() { // 数字输出装置。
		for n := range numberChan3 { // 不断的从数字通道3中接收数字，直到该通道关闭。
			fmt.Println(n) // 打印数字。
		}
		waitGroup.Done() // 表示一个操作完成。
	}()

	for i := 0; i < 100; i++ { // 先后向数字通道1传送100个范围在[0,100)的随机数。
		numberChan1 <- rand.Int63n(100)
	}
	close(numberChan1) // 数字发送完毕，关闭数字通道1。
	waitGroup.Wait()   // 等待前面那组操作（共3个）的完成。
}
