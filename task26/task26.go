package task26

import (
	"fmt"
	"unicode"
)

func TestUniq(str string) bool {
	uniqSymbols := map[rune]struct{}{}

	for _, val := range str {
		lowered := unicode.ToLower(val)
		if _, ok := uniqSymbols[lowered]; !ok {
			uniqSymbols[lowered] = struct{}{}
		} else {
			return false
		}
	}

	return true
}

func Task26() {
	// The first 'a' is russian letter in the both cases

	fmt.Println(TestUniq("абвгдA")) // The last 'A' - english letter
	fmt.Println(TestUniq("абвгдА")) // The last 'A' - russian letter
}
