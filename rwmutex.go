package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func tryRWMutex(kkkk int) {
	if kkkk == 0 {
		return
	}
	tryAnother()
	fmt.Printf("=======\n\n\n")
	lock := &sync.RWMutex{}
	lock.Lock()

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("the %vth goroutine prepare to start...\n", i)
			wg.Done()
			lock.RLock()
			//wg.Done() //wg.done here cause deadlock
			fmt.Printf("the %vth goroutine got read mutex, after 1 second sleep, unlock the locker\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}
	wg.Wait()

	fmt.Println("prepare unlock the wmutex. rmutex stop deadloop")
	lock.Unlock()

	lock.Lock()
	fmt.Println("program exits")
	lock.Unlock()
}

func tryAnother() {
	var wg sync.WaitGroup
	var rw sync.RWMutex
	var read func(int)
	var write func(int)
	var count int

	r := rand.New(rand.NewSource(time.Now().Unix()))

	read = func(n int) {
		rw.RLock() //读锁可以多次上锁
		fmt.Printf("read goroutine %d is reading...\n", n)

		v := count

		fmt.Printf("read goroutine %d finished reading, value is %d\n", n, v)
		wg.Done()
		rw.RUnlock()
	}

	write = func(n int) {
		rw.Lock()
		fmt.Printf("write goroutine %d is writting...\n", n)
		v := r.Intn(10000)

		count = v

		fmt.Printf("write goroutine %d finished writting, value is %d\n", n, v)
		wg.Done()
		rw.Unlock()
	}

	wg.Add(10)

	for i := 0; i < 5; i++ {
		go read(i)
	}

	for i := 0; i < 5; i++ {
		go write(i)
	}

	wg.Wait()
}
