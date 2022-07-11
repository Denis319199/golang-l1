package task19

import "fmt"

func ReverseString(str string) string {
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length>>1; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

func Task19() {
	str := "Привет мир!"
	fmt.Println(ReverseString(str))
	fmt.Println(str)
}
