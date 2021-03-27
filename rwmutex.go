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

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("the %vth goroutine prepare to start...\n", i)
			wg.Done()
			lock.RLock()
			//wg.Done()
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
