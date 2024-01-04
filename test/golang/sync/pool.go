package main

import (
	"fmt"
	"sync"
)

// sync.Pool 是在sync包下的内存池组件, 用来实现对象的复用,以避免操作频繁的内存分配和gc
// sync.Pool 不会永久保存对象,pool中的对象在一定的时间内会被gc回收
// 所以sync.Pool不支持持久化对象
// 其本身的并发安全的，支持多个goroutine并发的向pool存取数据

type Student struct {
	Name string
	Age  int
}

func PoolMain() {
	pool := sync.Pool{
		New: func() any {
			return &Student{
				Name: "zhangsan",
				Age:  18,
			}
		},
	}

	st := pool.Get().(*Student)
	println(st.Name, st.Age)
	fmt.Printf("addr is %p \n", st)

	pool.Put(st)

	st = pool.Get().(*Student)
	println(st.Name, st.Age)
	fmt.Printf("addr is %p \n", st)
}
