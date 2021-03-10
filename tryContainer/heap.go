package main

type heapNode struct {
	name string
	num int
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

func (h *heapSet) Push() {

}