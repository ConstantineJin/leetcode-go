package main

import "math/bits"

func sumIndicesWithKSetBits(nums []int, k int) (ans int) {
	for i, num := range nums {
		if bits.OnesCount(uint(i)) == k {
			ans += num
		}
	}
	return
}
