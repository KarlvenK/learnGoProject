package sync_once

import (
	"fmt"
	"sync"
)

type Obj struct{}

var one sync.Once
var instance *Obj

func New() *Obj {
	one.Do(func() {
		fmt.Println("Start creating Obj...")
		instance = &Obj{}
		fmt.Println("Obj created.")
	})
	return instance
}

func Once() {
	fmt.Println("Begin\nCall func New() for 10 times...")
	for i := 0; i < 10; i++ {
		_ = New()
	}
	fmt.Println("End.")
}
