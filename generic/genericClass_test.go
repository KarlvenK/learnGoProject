package generic

import (
	"fmt"
	"testing"
)

func TestGenericClass(t *testing.T) {
	f := func(x int) byte {
		return byte(int('0') + x)
	}
	nums := []int{1, 2, 3, 4, 5, 6}
	ret := Map(nums, f)
	fmt.Println(ret)

	ff := func (a byte, b int) byte {
		return byte(int(a) + b)
	}

	fmt.Println(Reduce(nums, '0', ff))

	fmt.Println(GenericSum(1,2,3,4,5,6))
	fxx()
}

func TestSort(t *testing.T) {
	nums := []int{5,1,5,7,235,12,4123,63,45,34651,234}
	Sort(nums, func(t1, t2 int) bool {
		return t1 < t2
	})
	fmt.Println(nums)
	fmt.Println(Pick(nums))
	doFoo()
	/*
	switch interface{}(1).(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("err")
	}
	*/
	fmt.Println(Equal(1, 2))	
}