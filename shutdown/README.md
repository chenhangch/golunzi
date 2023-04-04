# 实现一个优雅的关闭功能
即在应用程序关闭时，能够优雅的处理正在执行的任务和连接，而不是突然中断它们

代码实现集中在`shutdown.go`中

1. **ShutdownCallback** 接口和 **ShutdownFunc** 类型：用于定义和实现 ShutdownCallbacks，即在关闭时需要执行的回调函数。

2. **ShutdownManager** 接口：定义了 ShutdownManager 的行为，包括获取名称、启动 ShutdownManager、开始关闭、结束关闭等。

3. **ErrorHandler** 接口和 **ErrorFunc** 类型：用于定义和实现处理异步错误的函数。

4. **GSInterface** 接口：定义了 GracefulShutdown 结构体的行为，包括启动 ShutdownManager、报告错误、添加 ShutdownCallback 等。

5. **GracefulShutdown** 结构体：实现了上述接口和类型，并提供了 Start、AddShutdownManager、AddShutdownCallback、SetErrorHandler、StartShutdown、ReportError 等方法，用于启动和管理 ShutdownManager 和 ShutdownCallbacks。

在`posixsignal.go`中，实现了ShutdownManager 接口，使用POSIX信号实现优雅关闭

通过通道`sign`，接收操作系统的信号，从而调用**gs.StartShutdown(psm)**启动优雅关闭
```go
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGNAL...)
```
