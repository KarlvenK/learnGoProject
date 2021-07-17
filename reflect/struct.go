package reflect

import (
	"fmt"
	"reflect"
)

func Try() {
	type node struct {
		name string
		age  int
	}

	var item = node{
		name: "a",
		age:  1,
	}

	t := reflect.TypeOf(item)
	v := reflect.ValueOf(item)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println(v.Kind())
}
