package main

import (
	"math/bits"
	"slices"
)

func maxGoodNumber(nums []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int { return (b<<bits.Len(uint(a)) | a) - (a<<bits.Len(uint(b)) | b) })
	for _, num := range nums {
		ans = ans<<bits.Len(uint(num)) | num
	}
	return
}
