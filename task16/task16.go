package task16

import "fmt"

func QuickSort(values []int) {
	end := len(values) - 1
	if end <= 0 {
		return
	}

	left, right := 0, end
	pivot := values[right/2]

	for left <= right {
		for values[left] < pivot {
			left++
		}

		for values[right] > pivot {
			right--
		}

		if left < right {
			values[left], values[right] = values[right], values[left]
			left++
			right--
		} else if left == right {
			left++
			right--
		}
	}

	if right != 0 {
		QuickSort(values[:right+1])
	}
	if left != end {
		QuickSort(values[left:])
	}
}

func Task16() {
	arr := []int{3, 2, 1, 4, 5, 1}
	QuickSort(arr)
	fmt.Println(arr)
}
