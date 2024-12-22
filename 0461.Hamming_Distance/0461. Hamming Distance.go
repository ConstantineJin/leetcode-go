package main

import "math/bits"

func hammingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
