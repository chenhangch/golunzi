package main

import (
	"fmt"
	"time"
)

// select 是go语言底层提供的一种多路复用模型

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func() {
		select {
		case v := <-ch1:
			fmt.Printf("Received from ch1, val = %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received from ch1, val = %d\n", v)
		default:
			fmt.Println("default !!!")
		}
		time.Sleep(time.Second)
	}()
	ch1 <- 1
	time.Sleep(time.Second)
	ch2 <- 2
	time.Sleep(4 * time.Second)
}
