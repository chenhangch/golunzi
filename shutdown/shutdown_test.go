package shutdown

import "testing"


type SMShutdownStartFunc func() error

func (f SMShutdownStartFunc) GetName() string {
	return "test-sm"
}

func (f SMShutdownStartFunc) ShutdownStart() error {
	return f()
}

func (f SMShutdownStartFunc) ShutdownFinish() error {
	return nil
}

func (f SMShutdownStartFunc) Start(gs GSInterface) error {
	return nil
}

type SMFinishFunc func() error

func (f SMFinishFunc) GetName() string {
	return "test-sm"
}

func (f SMFinishFunc) ShutdownStart() error {
	return nil
}

func (f SMFinishFunc) ShutdownFinish() error {
	return f()
}

func (f SMFinishFunc) Start(gs GSInterface) error {
	return nil
}

type SMStartFunc func() error

func (f SMStartFunc) GetName() string {
	return "test-sm"
}

func (f SMStartFunc) ShutdownStart() error {
	return nil
}

func (f SMStartFunc) ShutdownFinish() error {
	return nil
}

func (f SMStartFunc) Start(gs GSInterface) error {
	return f()
}

func TestCallbacksGetCalled(t *testing.T) {
	gs := New()

	c := make(chan int, 100)
	for i := 0; i < 15; i++ {
		gs.AddShutdownCallback(ShutdownFunc(func(s string) error {
			c <- 1
			return nil
		}))
	}

	gs.StartShutdown(SMFinishFunc(func() error {
		return nil
	}))

	if len(c) != 15 {
		t.Error("Expected 15 elements in channel, got ", len(c))
	}
}