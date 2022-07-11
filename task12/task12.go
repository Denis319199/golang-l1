package task12

import "fmt"

func CreateSet[T comparable](values []T) map[T]struct{} {
	m := make(map[T]struct{})

	for _, val := range values {
		m[val] = struct{}{}
	}

	return m
}

func Task12() {
	// I am not sure that I understood the task correctly
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(CreateSet(seq))
}
