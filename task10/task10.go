package task10

import (
	"math/rand"
)

func GroupBy[T any, G comparable](values []T, f func(T) G) map[G][]T {
	m := map[G][]T{}

	for _, val := range values {
		key := f(val)
		arr, ok := m[key]
		if ok {
			m[key] = append(arr, val)
		} else {
			m[key] = []T{val}
		}
	}

	return m
}

func MergeGroupByResults[T any, G comparable](f, s map[G][]T) map[G][]T {
	if len(f) < len(s) {
		for key, fArr := range f {
			sArr, ok := s[key]
			if ok {
				s[key] = append(sArr, fArr...)
			} else {
				s[key] = fArr
			}
		}

		return s
	}

	for key, sArr := range s {
		fArr, ok := f[key]
		if ok {
			f[key] = append(fArr, sArr...)
		} else {
			f[key] = sArr
		}
	}

	return f
}

func SplitIntoTasks[T, R any](values []T, op func([]T) R, merge func(R, R) R, workers int) R {
	middleNum := workers >> 1
	end := len(values)

	if middleNum != 0 {
		var middle int
		if workers&1 == 0 {
			middle = end >> 1
		} else {
			middle = end / workers * middleNum
		}

		channel := make(chan R)
		go func() {
			channel <- SplitIntoTasks(values[:middle], op, merge, middleNum)
			close(channel)
		}()

		val := SplitIntoTasks(values[middle:], op, merge, workers-middleNum)
		return merge(val, <-channel)
	}

	return op(values)
}

func GroupByConcurrent[T any, G comparable](values []T, f func(T) G) map[G][]T {
	workersCount := len(values) / 250000
	if workersCount > 8 {
		workersCount = 8
	}

	if workersCount != 0 {
		wrapper := func(values []T) map[G][]T {
			return GroupBy(values, f)
		}

		return SplitIntoTasks(values, wrapper, MergeGroupByResults[T, G], workersCount)
	}

	return GroupBy(values, f)
}

func GetTemperatureGroup(t float32) int {
	rounded := int(t)
	return rounded - rounded%10
}

func Task10() {
	var arr [1000000]float32
	for i := 0; i < 1000000; i++ {
		arr[i] = rand.Float32()
	}

	GroupByConcurrent(arr[:], GetTemperatureGroup)
}
