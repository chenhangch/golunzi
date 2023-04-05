package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 获取当前文件的绝对路径
	absPath, err := filepath.Abs("./README.md")
	if err != nil {
		fmt.Println(err)
	}

	// 获取文件所在的目录
	dir := filepath.Dir(absPath)

	fmt.Println("文件的绝对路径：", absPath)
	fmt.Println("文件所在的目录：", dir)
}