package main

import "math/bits"

func minOperations(nums []int, k int) int {
	for _, num := range nums {
		k ^= num
	}
	return bits.OnesCount(uint(k))
}
