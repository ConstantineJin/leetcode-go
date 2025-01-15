package main

import "math/bits"

func minimizeXor(num1, num2 int) int {
	one1, one2 := bits.OnesCount(uint(num1)), bits.OnesCount(uint(num2))
	for ; one1 > one2; one2++ {
		num1 &= num1 - 1
	}
	for ; one1 < one2; one2-- {
		num1 |= num1 + 1
	}
	return num1
}
