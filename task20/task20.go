package task20

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWordsComplex(str string) string {
	// Tries to find words manually
	runes := []rune(str)
	length := len(runes)

	index := 0
	start, end := 0, 0

	for index < length {
		for index < length && unicode.IsSpace(runes[index]) {
			start++
			index++
		}
		end = start + 1
		index++

		for index < length && !unicode.IsSpace(runes[index]) {
			end++
			index++
		}
		index++

		for i := 0; i < (end-start)>>1; i++ {
			runes[start+i], runes[end-1-i] = runes[end-1-i], runes[start+i]
		}

		start = end + 1
	}

	return string(runes)
}

// ReverseWords - works only for words separated by only one space
func ReverseWords(str string) string {
	builder := strings.Builder{}
	builder.Grow(len(str))

	for _, word := range strings.Split(str, " ") {
		runes := []rune(word)
		for i := len(runes) - 1; i >= 0; i-- {
			builder.WriteRune(runes[i])
		}

		builder.WriteString(" ")
	}

	return builder.String()
}

func Task20() {
	str := "Привет, Hello мир, world!"
	fmt.Println(ReverseWordsComplex(str))
	fmt.Println(ReverseWords(str))
	fmt.Println(str)

	// Works for complex strings
	fmt.Println(ReverseWordsComplex(`МИР
ТРУД
МАЙ  !1`))
}
