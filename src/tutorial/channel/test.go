/*
sync.WaitGroup的使用也非常简单，先是使用Add 方法设设置计算器为2，每一个goroutine的函数执行完之后，就调用Done方法减1。Wait方法的意思是如果计数器大于0，就会阻塞，所以main 函数会一直等待2个goroutine完成后，再结束。
*/
/*
对于逻辑处理器的个数，不是越多越好，要根据电脑的实际物理核数，如果不是多核的，设置再多的逻辑处理器个数也没用，如果需要设置的话，一般我们采用如下代码设置。
runtime.GOMAXPROCS(runtime.NumCPU())
*/
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
		defer waitGroup.Done()
		for n := range numberChan1 {
			if n%2 == 0 {
				numberChan2 <- n
			} else {
				fmt.Printf("Filter %d. [filter 1]\n", n)
			}
		}
		close(numberChan2)
		// waitGroup.Done()
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
