package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 协程池
// 简单理解为就是一个池子里面存放固定数量的共routine
// 每当有一个任务来的时候，会将这个池子里面一个空闲的goroutine去处理，如果没有空闲的goroutine，任务就会阻塞
// 协程池有三个角色 Worker、Task、Pool

type Task struct {
	f func() error // 具体执行逻辑
}

func NewTask(funcArg func() error) *Task {
	return &Task{
		f: funcArg,
	}
}

type Pool struct {
	RunningWorkers int64
	Capacity       int64
	JobCh          chan *Task
	sync.Mutex
}

func NewPool(capacity int64, taskNum int) *Pool {
	return &Pool{
		Capacity: capacity,
		JobCh:    make(chan *Task, taskNum),
	}
}

func (p *Pool) GetCap() int64 {
	return p.Capacity
}

func (p *Pool) incRunning() {
	atomic.AddInt64(&p.RunningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt64(&p.RunningWorkers, -1)
}

func (p *Pool) GetRunningWorkers() int64 {
	return atomic.LoadInt64(&p.RunningWorkers)
}

func (p *Pool) run() {
	p.incRunning()
	go func() {
		defer func() {
			p.decRunning()
		}()
		for task := range p.JobCh {
			task.f()
		}
	}()
}

func (p *Pool) AddTask(task *Task) {
	p.Lock()
	defer p.Unlock()
	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}

	p.JobCh <- task
}

func main() {
	pool := NewPool(3, 10)
	for i := 0; i < 20; i++ {
		pool.AddTask(NewTask(func() error {
			fmt.Printf("I am Task\n")
			return nil
		}))
	}

	time.Sleep(1e9)
}
