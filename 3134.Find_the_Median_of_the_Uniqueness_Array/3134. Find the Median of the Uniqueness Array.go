package main

import "sort"

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	k := (n*(n+1)/2 + 1) / 2                       // 中位数就是第 k 小的值
	return sort.Search(n-1, func(upper int) bool { // 二分查找 distinct 值不超过 upper 的子数组的个数
		upper++
		var cnt, left int // 满足要求的子数组个数为 cnt
		freq := make(map[int]int)
		for right, in := range nums { // 滑动窗口
			freq[in]++
			for len(freq) > upper {
				out := nums[left]
				freq[out]--
				if freq[out] == 0 {
					delete(freq, out)
				}
				left++
			}
			cnt += right - left + 1
			if cnt >= k {
				return true
			}
		}
		return false
	}) + 1
}
