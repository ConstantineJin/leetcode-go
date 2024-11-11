package main

import "math"

type pair struct{ or, left int }

func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	var ors []pair // 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
	for i, num := range nums {
		ors = append(ors, pair{0, i})
		var j int
		for idx := range ors {
			p := &ors[idx]
			p.or |= num
			if p.or >= k {
				ans = min(ans, i-p.left+1)
			}
			if ors[j].or == p.or {
				ors[j].left = p.left // 原地去重：合并相同值，左端点取靠右的
			} else {
				j++
				ors[j] = *p
			}
		}
		ors = ors[:j+1] // 去重：移除多余元素
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
