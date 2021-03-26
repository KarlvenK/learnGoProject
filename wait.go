package main

import (
	"fmt"
	"sync"
)

func tryWait(kkkk int) {
	if kkkk == 0 {
		return
	}
	// 用信道来标记完成
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done

	fmt.Println("============")

	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)
	wg.Wait()
}

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}
