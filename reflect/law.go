package reflect

import (
	"fmt"
	"reflect"
)

func InterfaceToObject() { //law1
	var author = "Karlven"
	fmt.Println("var author string = \"Karlven\"")
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
}

func ObjectToInterface() { //law2
	v := reflect.ValueOf(1)
	fmt.Println(v.Interface().(int))
}

func Law3() {
	/* wrong way
	i := 1
	v := reflect.ValueOf(i)
	v.SetInt(10)
	fmt.Println(i)
	*/

	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
}
