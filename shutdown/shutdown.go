// learn from giuhub.com/marmotedu/iam/pkg/shutdown
package shutdown

import "sync"

// ShutdownCallback 提供了shutdown 前后需要执行的回调函数
type ShutdownCallback interface {
	OnShutdown(string) error
}

// ShutdownFunc是一个helper类型，所以你可以很容易地提供匿名函数ShutdownCallbacks。
type ShutdownFunc func(string) error

// OnShutdown 定义触发关机时需要运行的操作。
func (f ShutdownFunc) OnShutdown(shutdownManager string) error {
	return f(shutdownManager)
}

// ShutdownManager 接口：定义了 ShutdownManager 的行为，包括获取名称、启动 ShutdownManager、开始关闭、结束关闭等
type ShutdownManager interface {
	GetName() string
	Start(gs GSInterface) error
	ShutdownStart() error
	ShutdownFinish() error
}

// ErrorHandler 接口和 ErrorFunc 类型：用于定义和实现处理异步错误的函数
type ErrorHandler interface {
	OnError(err error)
}

type ErrorFunc func(err error)

// OnError 定义发生错误时运行所需的操作
func (f ErrorFunc) OnError(err error) {
	f(err)
}

// GSInterface 接口：定义了 GracefulShutdown 结构体的行为，包括启动 ShutdownManager、报告错误、添加 ShutdownCallback 等
type GSInterface interface {
	StartShutdown(sm ShutdownManager)
	ReportError(err error)
	AddShutdownCallback(shutdownCallback ShutdownCallback)
}

// GracefulShutdown 结构体：实现了上述接口和类型，并提供了 
// Start、AddShutdownManager、AddShutdownCallback、
// SetErrorHandler、StartShutdown、ReportError 等方法，
// 用于启动和管理 ShutdownManager 和 ShutdownCallbacks
type GracefulShutdown struct {
	callbacks []ShutdownCallback
	managers  []ShutdownManager
	errorHandler ErrorHandler
}

// New 初始化GracefulShutdown
func New() *GracefulShutdown {
	return &GracefulShutdown{
		callbacks: make([]ShutdownCallback, 0,10),
		managers: make([]ShutdownManager, 0, 3),
	}
}

func (gs *GracefulShutdown) Start() error {
	for _, manager := range gs.managers {
		if err := manager.Start(gs) ;err != nil {
			return err
		}
	}
	return nil
}

func (gs *GracefulShutdown) AddShutdownManager(manager ShutdownManager) {
	gs.managers = append(gs.managers, manager)
}

func (gs *GracefulShutdown) AddShutdownCallback(shutdownCallback ShutdownCallback)  {
	gs.callbacks = append(gs.callbacks, shutdownCallback)
}

func (gs *GracefulShutdown) SetErrorHandler(errorHandler ErrorHandler){
	gs.errorHandler = errorHandler
}

// StartShutdown从ShutdownManager调用，将启动shutdown。
// 首先在Shutdownmanager上调用ShutdownStart，
// 调用所有的ShutdownCallbacks，等待回调完成，
// 并在ShutdownManager上调用ShutdownFinish。
func (gs *GracefulShutdown) StartShutdown(sm ShutdownManager) {
	gs.ReportError(sm.ShutdownStart())

	var wg sync.WaitGroup
	for _, shutdownCallback := range gs.callbacks {
		wg.Add(1)
		go func(shutdownCallback ShutdownCallback) {
			defer wg.Done()

			gs.ReportError(shutdownCallback.OnShutdown(sm.GetName()))
		}(shutdownCallback)
	}
	wg.Wait()

	gs.ReportError(sm.ShutdownFinish())
}

// ReportError是一个可以用来向ErrorHandler报告错误的函数。它在ShutdownManagers中使用。
func (gs *GracefulShutdown) ReportError(err error) {
	if err != nil && gs.errorHandler != nil{
		gs.errorHandler.OnError(err)
	}
}