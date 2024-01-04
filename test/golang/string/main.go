package main

import (
	"fmt"
	"unsafe"
)

// 字符串是所有8bit字节的集合，但不一定是UTF-8编码的文本
// 字符串可以是empty，但不能为nil
// empty字符串就是一个没有任何字符串的空串""
// 字符串不可以被修改，所以字符串类型的值是不可变的
var text string

// stringStruct string数据结构
// stringStruct是字符串在运行状态下的表现，创建一个string的时候，可以分为两步
// 1. 根据给定的字符创建出stringStruct结构
// 2. 将stringStruct结构转化为string类型
type stringStruct struct {
	str unsafe.Pointer // str指向字符串的首地址
	len int            // 字符串的长度, len字段存储的实际字节数而不是字符数，对于非单字节编码的字符，其结果可能多于字符个数
}

func main() {
	text = "Hello"
	for _, v := range text {
		fmt.Printf("%d \n", v)
	}
	
	strByte := []byte(text)
	strByte[2] = 'A'
	fmt.Println(string(strByte))
}
