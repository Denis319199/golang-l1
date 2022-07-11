package task17

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func BinarySearch(values []int, s int) (int, error) {
	start, end := 0, len(values)-1
	index := end >> 1

	for start <= end {
		val := values[index]

		if val == s {
			return index, nil
		} else if val < s {
			start = index + 1
		} else {
			end = index - 1
		}
		index = (start + end) >> 1
	}

	return -1, ErrNotFound
}

func Task17() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	_, err := BinarySearch(arr, 11)

	fmt.Println(err)

}
