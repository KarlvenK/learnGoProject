package main

import (
	"fmt"
	"reflect"
)

type human struct {
	num  string
	age  int
	gene bool
}

func (h *human) weight() int {
	return 1
}

func (h *human) length() int {
	return len(h.num)
}

func (h *human) isMale() bool {
	return h.gene
}

type animal interface {
	weight() int
	length() int
	isMale() bool
}

func tryReflect(vv int) {
	if vv == 0 {
		return
	}
	var age interface{} = 25
	fmt.Printf("orginal interface type is %T, value is %v\n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	fmt.Printf("transform from interface to reflected object. Type : %T\n", t)
	fmt.Printf("~                                             Value : %T\n", v)

	kk := human{
		"kk",
		11,
		true,
	}

	i := animal(&kk)

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
}
