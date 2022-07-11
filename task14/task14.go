package task14

import (
	"fmt"
	"reflect"
)

func Foo(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.ValueOf(val).Kind() == reflect.Chan {
			fmt.Println("chan")
		}
	}
}

func Task14() {
	Foo(1.2)
	Foo(make(chan int))
	Foo(1)
	Foo(true)
	Foo("123")
}
