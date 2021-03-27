package main

import "fmt"

func tryChannel(v int) {
	if v == 0 {
		return
	}
	/*
		arr := []int {1, 2, 3, 4, 5, 6, 7}
		n := len(arr)
		c := make(chan int)
		go sum(arr[:n], c)
		go sum(arr[:n/2], c)
		x, y := <-c, <- c
		fmt.Println(x, y)

		fmt.Println("========")

		ch := make(chan int, 45)
		go putNum(cap(ch), ch)
		for i := range ch {
			fmt.Println(i)
		}
	*/
	testTree()
	fmt.Println("testTree end")
	pipeline := make(chan string)
	go func() {
		defer close(pipeline) //without it the goroutine will deadlock
		pipeline <- "helloworld"
		pipeline <- "hellochina"
	}()
	for data := range pipeline {
		fmt.Println(data)
	}
}

func testTree() {
	ch := make(chan int)
	t1 := treeNew(1)
	go Walk(t1, ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(t1, t1))
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	realWalk(t, ch)
	close(ch)
}
func realWalk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	realWalk(t.Left, ch)
	ch <- t.Value
	realWalk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

/*
func putNum(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x + y
	}
	close(ch)
}

/*
func sum(arr []int, c chan int) {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	c <- sum
}
*/
