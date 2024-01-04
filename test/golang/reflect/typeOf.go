package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func TypeOfMain() {
	var num int64 = 100
	ti := reflect.TypeOf(num)
	fmt.Println(ti.String())

	st := Student{
		Name: "zhangsan",
		Age:  18,
	}
	t2 := reflect.TypeOf(st)
	fmt.Println(t2.String())
}
