package task13

import "fmt"

func Task13() {
	// Option 1
	a, b := 666, 777
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

	// Option 2
	a, b = 666, 777
	fmt.Println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
