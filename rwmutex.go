package main

import (
	"fmt"
	"sync"
	"time"
)

func tryRWMutex(kkkk int) {
	if kkkk == 0 {
		return
	}

	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("the %vth goroutine prepare to start...\n", i)
			lock.RLock()
			fmt.Printf("the %vth goroutine got read mutex, after 1 second sleep, unlock the locker\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("prepare unlock the wmutex. rmutex stop deadloop")
	lock.Unlock()

	lock.Lock()
	fmt.Println("program exits")
	lock.Unlock()
}
