package task11

import "fmt"

func Intersect[T comparable](f, s map[T]struct{}) map[T]struct{} {
	res := map[T]struct{}{}

	if len(f) < len(s) {
		for key, _ := range f {
			if _, ok := s[key]; ok {
				res[key] = struct{}{}
			}
		}
	} else {
		for key, _ := range s {
			if _, ok := f[key]; ok {
				res[key] = struct{}{}
			}
		}
	}

	return res
}

func Task11() {
	fmt.Println(Intersect(
		map[int]struct{}{1: {}, 2: {}, 3: {}},
		map[int]struct{}{2: {}, 3: {}, 4: {}, 5: {}}))
}
