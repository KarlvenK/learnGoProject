package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var ErrInvalidPoolCap = errors.New("invalid pool cap")
var ErrPoolAlreadyClosed = errors.New("pool already closed")

const (
	RUNNING = 1
	STOPPED = 0
)

type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

type Pool struct {
	capacity       uint64
	runningWorkers uint64
	status         int64
	chTask         chan *Task
	sync.Mutex

	PanicHandler func(interface{})
}

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{
		capacity: capacity,
		status:   RUNNING,
		//initial the original tasks queue len(queue) = capacity
		chTask: make(chan *Task, capacity),
	}, nil
}

/* unsafe
func (p *Pool) run() {
	p.runningWorkers++

	go func() {
		defer func() {
			p.runningWorkers--
		}()

		for {
			select {
			case task, ok := <-p.chTask:
				if !ok{
					return
				}
				task.Handler(task.Params...)
			}
		}
	}()
}*/

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	// -n = ^(n - 1)
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return p.capacity
}

func (p *Pool) setStatus(status int64) bool {
	p.Lock()
	defer p.Unlock()

	if p.status == status {
		return false
	}

	p.status = status
	return true
}

func (p *Pool) run() {
	p.incRunning()
	go func() {
		defer func() {
			p.decRunning()
			if r := recover(); r != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Printf("worker panic: %s\n", r)
				}
			}
			p.checkWorker()
		}()

		for {
			select {
			case task, ok := <-p.chTask:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			}
		}
	}()
}

func (p *Pool) Put(task *Task) error {
	p.Lock()
	defer p.Unlock()

	if p.status == STOPPED {
		return ErrPoolAlreadyClosed
	}

	if p.GetRunningWorkers() < p.capacity {
		p.run()
	}

	if p.status == RUNNING {
		p.chTask <- task
	}
	return nil
}
func (p *Pool) checkWorker() {
	p.Lock()
	defer p.Unlock()
	if p.runningWorkers == 0 && len(p.chTask) > 0 {
		p.run()
	}
}

/*
func (p *Pool) Close() {
	p.setStatus(STOPPED)

	for len(p.chTask) > 0 {
		time.Sleep(1e6)
	}
	close(p.chTask) //it may cause panic for multiple close
					//so we need to lock it on
}*/
func (p *Pool) close() {
	p.setStatus(STOPPED)
	for len(p.chTask) > 0 {
		time.Sleep(1e6)
	}
	p.close()
}

func (p *Pool) Close() {
	p.Lock()
	defer p.Unlock()
	close(p.chTask)
}

func tryPool(kkkkkk int) {
	if kkkkkk == 0 {
		return
	}
	pool, err := NewPool(10)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 20; i++ {
		_ = pool.Put(&Task{
			Handler: func(v ...interface{}) {
				fmt.Println(v)
			},
			Params: []interface{}{i},
		})
	}
	time.Sleep(1e9)
}

//via https://segmentfault.com/a/1190000021468353
