package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	timeout := make(chan bool, 1)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			timeout <- true
			fmt.Println("协程设置完毕")
		}
	}()
	// for {
	select {
	case <-ch:
		fmt.Println("This is ch.")
	case <-timeout:
		fmt.Println("This is timeout.")
	}
	// }
}
