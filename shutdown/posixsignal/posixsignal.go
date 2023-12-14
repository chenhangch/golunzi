package posixsignal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chenhangch/golunzi/shutdown"
)

const Name = "PosixSignalManager"

type PosixSignalManager struct {
	signals []os.Signal
}

// NewPosixSignalManager 初始化PosixSignalManager，传入os.Signal，
// 默认值为：os.Interrupt | syscall.SIGTERM
func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = make([]os.Signal, 2)
		sig[0] = os.Interrupt
		sig[1] = syscall.SIGTERM
	}

	return &PosixSignalManager{
		signals: sig,
	}
}

func (psm *PosixSignalManager) GetName() string {
	return Name
}

func (psm *PosixSignalManager) Start(gs shutdown.GSInterface) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, psm.signals...)

		<-c

		gs.StartShutdown(psm)
	}()

	return nil
}

func (psm *PosixSignalManager) ShutdownStart() error {
	return nil
}

func (psm *PosixSignalManager) ShutdownFinish() error {
	os.Exit(0)

	return nil
}
