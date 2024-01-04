package main

import (
	"context"
	"fmt"
	"time"
)

// context 是一个标准库的接口
// type Context interface {
// 	...
// }

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
	
	// go Watch(ctx, "g1")
	// go Watch(ctx, "g2")

	// time.Sleep(time.Second * 12)
	// fmt.Println("end watching")

	// cancel()
	// time.Sleep(time.Second)

	ctx := context.WithValue(context.Background(), "name", "chenhang")
	go withValue(ctx)
	time.Sleep(time.Second)
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s exit \n", name)
			return
		default:
			fmt.Printf("%s watching... \n", name)
			time.Sleep(time.Second)
		}
	}
}
