package main

import (
	"math/bits"
	"slices"
)

// 暴力
//func findMaximumXOR(nums []int) (max int) {
//	n := len(nums)
//	xors := make([]int, 0)
//	for i := 0; i < n; i++ {
//		for j := i; j < n; j++ {
//			xors = append(xors, nums[i]^nums[j])
//		}
//	}
//	for _, xor := range xors {
//		if xor > max {
//			max = xor
//		}
//	}
//	return
//}

func findMaximumXOR(nums []int) (ans int) {
	highBit := bits.Len(uint(slices.Max(nums))) - 1
	seen := map[int]bool{}
	mask := 0
	for i := highBit; i >= 0; i-- { // 从最高位开始枚举
		clear(seen)
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		for _, x := range nums {
			x &= mask // 低于 i 的比特位置为 0
			if seen[newAns^x] {
				ans = newAns // 这个比特位可以是 1
				break
			}
			seen[x] = true
		}
	}
	return
}
