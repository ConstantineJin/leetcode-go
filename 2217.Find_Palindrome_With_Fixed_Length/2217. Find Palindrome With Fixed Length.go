package main

import "math"

func kthPalindrome(queries []int, intLength int) []int64 {
	var ans, base = make([]int64, len(queries)), int(math.Pow10((intLength - 1) / 2))
	for i, query := range queries {
		if query > 9*base {
			ans[i] = -1
			continue
		}
		var v = base + query - 1
		var x = v
		if intLength%2 == 1 {
			x /= 10
		}
		for ; x > 0; x /= 10 {
			v = v*10 + x%10
		}
		ans[i] = int64(v)
	}
	return ans
}
