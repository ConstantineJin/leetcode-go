package main

import "math/bits"

func minBitFlips(start, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}
