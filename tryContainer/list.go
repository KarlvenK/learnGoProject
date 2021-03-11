package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

func tryList() {
	l := list.New()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("input")
	for i := 0; i < 10; i++ {
		l.PushBack(r.Intn(100))
		fmt.Print(l.Back().Value, " ")
	}
	fmt.Println()
	fmt.Println("show")
	for auto := l.Front(); auto != nil; auto = auto.Next() {
		fmt.Print(auto.Value, " ")
	}
}
