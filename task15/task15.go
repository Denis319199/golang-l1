package task15

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var justString string

func someFunc() {
	// The problem is that we don't use the all length of the big string anymore,
	// but we retain a slice of this string, so the big one will not be collected
	// by GC hence extra memory is occupied
	v := createHugeString(1 << 10)

	// OPTION 1
	// Can use bytes.Buffer instead of strings.Builder
	var builder strings.Builder
	builder.WriteString(v[:100])
	justString = builder.String()

	// OPTION 2
	var bufff [100]byte
	copy(bufff[:], []byte(v[:100]))

	// OPTION 3
	justString = string([]byte(v[:100]))
}

func Task15() {
	someFunc()
	fmt.Println(justString)
}
