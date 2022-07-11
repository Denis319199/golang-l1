package task8

import "fmt"

func SetBit(value, bitNum int64, bitValue bool) int64 {
	if bitValue {
		return value | (1 << bitNum)
	}
	return value & (^(1 << bitNum))
}

func Task8() {
	var val int64
	val = SetBit(val, 10, true)
	val = SetBit(val, 8, true)
	val = SetBit(val, 10, false)
	fmt.Println(val)
}
