package main

import "math/bits"

func rangeBitwiseAnd(m, n int) int {
	return m &^ (1<<bits.Len(uint(m^n)) - 1)
}
