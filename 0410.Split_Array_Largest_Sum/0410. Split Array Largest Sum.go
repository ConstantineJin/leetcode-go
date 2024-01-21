package main

import "sort"

func splitArray(nums []int, k int) int {
	var sum, mx int
	for _, x := range nums {
		sum += x
		mx = max(mx, x)
	}
	var left, right = max(mx, (sum-1)/k+1), sum
	return left + sort.Search(right-left, func(mx int) bool {
		mx += left
		var cnt, s = 1, 0
		for _, x := range nums {
			if s+x <= mx {
				s += x
			} else { // 新划分一段
				if cnt == k {
					return false
				}
				cnt += 1
				s = x
			}
		}
		return true
	})
}
