package main

import (
	"fmt"

	"github.com/chenhangch/golunzi/test/golang.init/package1"
	"github.com/chenhangch/golunzi/test/golang.init/utils"
)

// go程序初始化过程：
// 包级别变量的初始化优先于包init函数的执行
// 一个包下面可以有多个init函数，每个文件也可以有多个init函数
// 多个init函数按照它们的文件名顺序逐个初始化
// 应用初始化时初始化工作的顺序时，从被导入的最深层包开始进行初始化，层层递出最后到main包
// 不管包被导入多少次，包内的init函数只会执行一次
// 应用在所有初始化工作完成后才会执行main函数


func init() {
	fmt.Println("init func1 in main")
}

func init() {
	fmt.Println("init func2 in main")
}

// var MainValue1 = utils.TraceLog("init M_v1", package1.V1+10)
// var MainValue2 = utils.TraceLog("init M_v2", package1.V2+10)

func main() {
	utils.TraceLog("init M_v1", package1.V1+10)
	fmt.Println("main func in main")
}
