package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic 和 mutex 的区别
// 通常 mutex 用于保护一段执行逻辑, atomic 主要是对变量进行操作
// 底层实现上: mutex有操作系统调度器实现, atomic操作有底层硬件指令支持,保证CPU上执行不中断

func AtomicMain() {
	var sum int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&sum, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("sum is %v\n", sum)
}
