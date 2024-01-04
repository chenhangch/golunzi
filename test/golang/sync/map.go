package main
import (
	"fmt"
	"strconv"
	"sync"
)

// go内置的Map并不是并发安全，在多个goroutine同时操作时，会有并发问题
// 但是sync.Map没有提供获取map数量的方法，需要我们自己对sync.Map进行遍历自行计算
// sync.Map为了保证并发安全会有一些性能损耗

var m = make(map[string]int, 1)
var mu sync.Mutex

func getVal(key string) int {
	return m[key]
}

func setVal(key string, value int) {
	m[key] = value
}

func MapMain() {
	wg := sync.WaitGroup{}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func(num int) {
			defer func() {
				wg.Done()
				mu.Unlock()
			}()
			key := strconv.Itoa(num)
			mu.Lock()
			setVal(key, num)
			fmt.Printf("key= :%v, val:= %v\n", key, getVal(key))
		}(i)
	}
	wg.Wait()

	fmt.Println(">==== use sync.Map example ====< ")
	fmt.Println(">==== use sync.Map example ====< ")
	fmt.Println(">==== use sync.Map example ====< ")

	var sm sync.Map
	// 写入
	sm.Store("name", "chenhang")
	sm.Store("age", "18")

	// 读取
	age, _ := sm.Load("age")
	fmt.Println(age)

	sm.Range(func(key, value any) bool {
		fmt.Printf("key is %v, val is %v \n", key, value)
		return true
	})

	sm.Delete("age")
	age, ok := sm.Load("age")
	fmt.Println(age, ok)

	sm.LoadOrStore("name", "zhangsan")
	name, _ := sm.Load("name")
	fmt.Println(name)
}
