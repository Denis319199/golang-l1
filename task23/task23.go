package task23

import "fmt"

func RemoveComplex[T any](arr []T, index int) []T {
	length := len(arr)

	if index >= length || index < 0 {
		return arr
	}

	if index+1 < length {
		copy(arr[index:], arr[index+1:])
	}

	return arr[:length-1]
}

func Remove[T any](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}

func Task23() {
	val := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 10; i >= 0; i-- {
		val = Remove(val, i)
		fmt.Println(val)
	}

	val = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 10; i >= 0; i-- {
		val = RemoveComplex(val, i)
		fmt.Println(val)
	}
}
