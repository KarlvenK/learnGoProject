package main

import (
	"fmt"
	"time"
)

func tryRoutine(v int) {
	if v == 0 {
		return
	}

	var flag int
	flag = 0

	if flag == 0 {
		go spinner(100 * time.Millisecond)
		const n = 45
		fibN := fib(n)
		fmt.Printf("\rFibonacci(%d) = %d \n", n, fibN)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range  "_\\|/" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay * 2)
		}
	}
}

// fib() runs slowly
func fib(x int) int {
	if x < 2 {
		return 1
	}
	return  fib(x - 1) + fib(x - 2)
}