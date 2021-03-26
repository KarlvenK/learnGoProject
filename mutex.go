package main

import (
	"fmt"
	"sync"
)

func tryMutex(v int) {
	if v == 0 {
		return
	}

	fmt.Printf("use waitGroup without mutex\n")

	var wg sync.WaitGroup
	count := 0

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go addd(&count, &wg)
	}

	wg.Wait()
	fmt.Printf("count equals to %d\n", count)
	fmt.Printf("===========\n")

	fmt.Printf("use mutex\n")

	var lock *sync.Mutex

	lock = new(sync.Mutex)
	// lcok := &sync.Mutex{}
	wg.Add(10)
	count = 0
	for i := 0; i < 10; i++ {
		go adddd(&count, &wg, lock)
	}
	wg.Wait()
	fmt.Printf("count equals to %d\n", count)
}

func addd(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count = *count + 1
	}
	wg.Done()
}

func adddd(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}
