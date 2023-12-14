package main

import (
	"fmt"
	"time"

	"github.com/chenhangch/golunzi/shutdown"
	"github.com/chenhangch/golunzi/shutdown/posixsignal"
)

func main() {
	// 初始化shutdown实例
	gs := shutdown.New()
	// 添加 posix shutdown 管理器
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())
	// 添加自定义优雅关机逻辑
	gs.AddShutdownCallback(shutdown.ShutdownFunc(func(s string) error {
		fmt.Println("Shutdown callback start")
		time.Sleep(time.Second * 2)
		fmt.Println("Shutdown callback finished")
		return nil
	}))
	// 启动容器
	if err := gs.Start(); err != nil {
		fmt.Println("Start: ", err)
		return
	}
	// 其它程序代码
	time.Sleep(time.Hour)
}
