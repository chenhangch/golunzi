package gracefulshutdown

import (
    "context"
    "fmt"
    "sync"
)

// ShutdownCallback is the interface that defines the callback function
// that will be executed when the application is shutting down.
type ShutdownCallback interface {
    Execute() error
}

// ShutdownFunc is a type that wraps a function as a ShutdownCallback.
type ShutdownFunc func() error

// Execute calls the wrapped function as a ShutdownCallback.
func (sf ShutdownFunc) Execute() error {
    return sf()
}

// ShutdownManager is the interface that defines the behavior of a
// shutdown manager, which is responsible for managing the shutdown
// process of the application.
type ShutdownManager interface {
    GetName() string
    Start() error
    BeginShutdown()
    EndShutdown()
}

// ErrorHandler is the interface that defines the callback function
// that will be executed when an error occurs during the shutdown
// process.
type ErrorHandler interface {
    HandleError(error)
}

// ErrorFunc is a type that wraps a function as an ErrorHandler.
type ErrorFunc func(error)

// HandleError calls the wrapped function as an ErrorHandler.
func (ef ErrorFunc) HandleError(err error) {
    ef(err)
}

// GSInterface is the interface that defines the behavior of the
// GracefulShutdown struct, which is responsible for managing the
// shutdown process of the application.
type GSInterface interface {
    Start(context.Context) error
    AddShutdownManager(ShutdownManager)
    AddShutdownCallback(ShutdownCallback)
    SetErrorHandler(ErrorHandler)
    StartShutdown()
    ReportError(error)
}

// GracefulShutdown is a struct that implements the GSInterface.
type GracefulShutdown struct {
    name            string
    shutdownManagers []ShutdownManager
    shutdownMutex   sync.Mutex
    shutdownStarted bool
    shutdownCtx     context.Context
    shutdownCancel  context.CancelFunc
    callbacks       []ShutdownCallback
    errorHandler    ErrorHandler
}

// NewGracefulShutdown creates a new GracefulShutdown with the given name.
func NewGracefulShutdown(name string) *GracefulShutdown {
    return &GracefulShutdown{
        name: name,
    }
}

// Start starts the GracefulShutdown and returns an error if any.
func (gs *GracefulShutdown) Start(ctx context.Context) error {
    gs.shutdownCtx, gs.shutdownCancel = context.WithCancel(ctx)
    for _, sm := range gs.shutdownManagers {
        if err := sm.Start(); err != nil {
            return fmt.Errorf("error starting shutdown manager: %s", err.Error())
        }
    }
    go gs.handleShutdown()
    return nil
}

// AddShutdownManager adds a ShutdownManager to the GracefulShutdown.
func (gs *GracefulShutdown) AddShutdownManager(sm ShutdownManager) {
    gs.shutdownMutex.Lock()
    defer gs.shutdownMutex.Unlock()
    gs.shutdownManagers = append(gs.shutdownManagers, sm)
}

// AddShutdownCallback adds a ShutdownCallback to the GracefulShutdown.
func (gs *GracefulShutdown) AddShutdownCallback(cb ShutdownCallback) {
    gs.shutdownMutex.Lock()
    defer gs.shutdownMutex.Unlock()
    gs.callbacks = append(gs.callbacks, cb)
}

// SetErrorHandler sets the ErrorHandler of the GracefulShutdown.
func (gs *GracefulShutdown) SetErrorHandler(eh ErrorHandler) {
    gs.errorHandler = eh
}

// StartShutdown begins the shutdown process of the GracefulShutdown.
func (gs *GracefulShutdown) StartShutdown() {
    gs.shutdownMutex.Lock()
    defer gs.shutdownMutex.Unlock()
    if gs.shutdownStarted {
        return
    }
    gs.shutdownStarted = true
    gs.shutdownCancel()
}

// ReportError reports an error to the GracefulShutdown.
func (gs *GracefulShutdown) ReportError(err error) {
    gs.errorHandler.HandleError(err)
}

// handleShutdown waits for the shutdown signal and executes the
// shutdown process of the application.
func (gs *GracefulShutdown) handleShutdown() {
    <-gs.shutdownCtx.Done()
    gs.shutdownMutex.Lock()
    defer gs.shutdownMutex.Unlock()
    for _, cb := range gs.callbacks {
        if err := cb.Execute(); err != nil {
            gs.errorHandler.HandleError(fmt.Errorf("error executing shutdown callback: %s", err.Error()))
        }
    }
    for _, sm := range gs.shutdownManagers {
        sm.BeginShutdown()
    }
    for _, sm := range gs.shutdownManagers {
        sm.EndShutdown()
    }
}

// GetName returns the name of the GracefulShutdown.
func (gs *GracefulShutdown) GetName() string {
    return gs.name
}