package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type heapNode struct {
	name string
	num  int
}

type heapSet []heapNode

func (h heapSet) Len() int {
	return len(h)
}

func (h heapSet) Less(i, j int) bool {
	return h[i].num < h[j].num
}

func (h heapSet) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapSet) Push(x interface{}) {
	*h = append(*h, x.(heapNode))
}
func (h *heapSet) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[:n-1]
	*h = old[:n-1]
	return x
}

func tryHeap() {
	var set heapSet
	r := rand.New(rand.NewSource(time.Now().Unix()))
	size := 15
	for i := 0; i < size; i++ {
		set = append(set, heapNode{strconv.Itoa(r.Intn(10000)), r.Intn(100)})
	}
	hPtr := &set
	heap.Init(hPtr)
	fmt.Println(*hPtr)
	for i := 0; i < size; i++ {
		if r.Int()&1 == 1 {
			auto := heapNode{strconv.Itoa(r.Intn(10000)), r.Intn(1000)}
			fmt.Printf("push %v\n", auto)
			heap.Push(hPtr, auto)
		} else {
			fmt.Printf("pop %v\n", (*hPtr)[0])
			heap.Pop(hPtr)
		}
	}
	fmt.Println(*hPtr)
}
