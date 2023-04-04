package main

import (
    "context"
    "fmt"
    "github.com/hangcodebug/golunzi/shutdown/cgracefulshutdown"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

// MyService is a simple service that needs to be gracefully shutdown.
type MyService struct {
    name string
	ctx  context.Context
}

// NewMyService creates a new MyService with the given name.
func NewMyService(name string) *MyService {
    return &MyService{
        name: name,
    }
}

// Start starts the MyService and blocks until the shutdown signal is received.
func (s *MyService) Start() error {
    fmt.Printf("Starting %s...\n", s.name)
    <-s.ctx.Done()
    fmt.Printf("Stopping %s...\n", s.name)
    return nil
}

func (s *MyService) GetName() string{
	return s.name
}

func (s *MyService) BeginShutdown() {
	fmt.Println("BeginShutdown")
}

func (s *MyService) EndShutdown() {
	fmt.Println("EndShutdown")
}

// MyShutdownManager is a simple shutdown manager that logs the
// shutdown process of the application.
type MyShutdownManager struct {
    name string
}

// NewMyShutdownManager creates a new MyShutdownManager with the given name.
func NewMyShutdownManager(name string) *MyShutdownManager {
    return &MyShutdownManager{
        name: name,
    }
}

// GetName returns the name of the MyShutdownManager.
func (m *MyShutdownManager) GetName() string {
    return m.name
}

// Start starts the MyShutdownManager and returns nil.
func (m *MyShutdownManager) Start() error {
    fmt.Printf("Starting %s...\n", m.name)
    return nil
}

// BeginShutdown logs the beginning of the shutdown process of the application.
func (m *MyShutdownManager) BeginShutdown() {
    fmt.Printf("Beginning shutdown of %s...\n", m.name)
}

// EndShutdown logs the end of the shutdown process of the application.
func (m *MyShutdownManager) EndShutdown() {
    fmt.Printf("Finished shutdown of %s...\n", m.name)
}

// main is the entry point of the application.
func main() {
    // Create a new GracefulShutdown.
    gs := gracefulshutdown.NewGracefulShutdown("MyApp")

    // Create a new MyService and add it to the GracefulShutdown.
    myService := NewMyService("MyService")
    gs.AddShutdownManager(myService)

    // Create a new MyShutdownManager and add it to the GracefulShutdown.
    myShutdownManager := NewMyShutdownManager("MyShutdownManager")
    gs.AddShutdownManager(myShutdownManager)

    // Add a ShutdownCallback that waits for some time before returning.
    callback1 := gracefulshutdown.ShutdownFunc(func() error {
        fmt.Println("Executing callback1...")
        time.Sleep(2 * time.Second)
        fmt.Println("Callback1 executed.")
        return nil
    })
    gs.AddShutdownCallback(callback1)

    // Add a ShutdownCallback that waits for some time before returning.
    callback2 := gracefulshutdown.ShutdownFunc(func() error {
        fmt.Println("Executing callback2...")
        time.Sleep(2 * time.Second)
        fmt.Println("Callback2 executed.")
        return nil
    })
    gs.AddShutdownCallback(callback2)

    // Set an ErrorHandler that logs the errors during the shutdown process.
    gs.SetErrorHandler(gracefulshutdown.ErrorFunc(func(err error) {
        fmt.Printf("Error during shutdown: %s\n", err.Error())
    }))

    // Start the GracefulShutdown.
    ctx, cancel := context.WithCancel(context.Background())
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := gs.Start(ctx); err != nil {
            fmt.Printf("Error starting GracefulShutdown: %s\n", err.Error())
        }
    }()

    // Wait for the interrupt signal and start the shutdown process.
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig
    fmt.Println("Received shutdown signal.")
    gs.StartShutdown()

    // Wait for the GracefulShutdown to complete.
    wg.Wait()
    cancel()
}