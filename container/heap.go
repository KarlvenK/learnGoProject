package container

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type intType []int

func (foo intType) Len() int {
	return len(foo)
}

//Less 大根堆
func (foo intType) Less(i, j int) bool {
	return foo[i] > foo[j]
}

func (foo intType) Swap(i, j int) {
	foo[i], foo[j] = foo[j], foo[i]
}

func (foo *intType) Push(x interface{}) {
	num := x.(int)
	*foo = append(*foo, num)
}

func (foo *intType) Pop() interface{} {
	num := (*foo)[len(*foo)-1]
	*foo = (*foo)[:len(*foo)-1]
	return num
}

func Try() {
	const tot = 1000
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < tot; i++ {
		nums = append(nums, rand.Intn(100000))
	}
	fmt.Println(nums)

	hp := make(intType, 0)
	heap.Init(&hp)

	for _, num := range nums {
		heap.Push(&hp, num)
	}

	for true {
		if hp.Len() == 0 {
			break
		}
		fmt.Println(heap.Pop(&hp))
	}
}
