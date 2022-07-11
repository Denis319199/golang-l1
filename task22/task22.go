package task22

import (
	"fmt"
	"math/big"
)

func Task22() {
	val, pow := big.NewInt(2), big.NewInt(100)
	val.Exp(val, pow, nil)
	fmt.Println(val.String())
}
