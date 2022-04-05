package generic

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Map[T1, T2 interface{}](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T1, T2 interface{}](s []T1, init T2, f func(T2, T1) T2) T2 {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

type wrapSort[T interface{}] struct {
	s []T
	cmp func(T, T) bool
}

func (s wrapSort[T]) Len() int {
	return len(s.s)
}
func (s wrapSort[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}
func (s wrapSort[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func Sort[T interface{}] (s []T, cmp func(T, T) bool) {
	sort.Sort(wrapSort[T]{s, cmp})
}

//Pick returns a randomly selected element in a given slice
//func Pick[S interface{~[]Elem}, Elem interface{}](s S) Elem
func Pick[S ~[]Elem, Elem any](s S) Elem {
	n := len(s)
	rand.Seed(time.Now().Unix())
	return s[rand.Intn(n)]
}

func GenericSum[S ~int](elems ...S) (sum S) {
	for _, v := range elems {
		sum += v
	}
	return
}

type S[T any] struct {
	a T
}

func fxx() {
	t := S[int] {
		a: 1,
	}
	fmt.Println(t)
}

type I[T any] interface {
	~int | ~int64 | ~int32
	M(v T) T
}

func Foo[T any]() {
	fmt.Println("do foo[T any] func")
}

func doFoo() {
	//x := Foo #error
	x := Foo[int] //ok
	x()
}

