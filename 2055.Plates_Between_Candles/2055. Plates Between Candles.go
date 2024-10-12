package main

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	left, right, prefixSums := make([]int, n), make([]int, n), make([]int, n+1)
	pre := -1
	for i, c := range s {
		prefixSums[i+1] = prefixSums[i]
		if c == '|' {
			pre = i
		} else {
			prefixSums[i+1]++
		}
		left[i] = pre
	}
	for i, pre := n-1, n; i >= 0; i-- {
		if s[i] == '|' {
			pre = i
		}
		right[i] = pre
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		l, r := right[query[0]], left[query[1]]
		if l < r {
			ans[i] = prefixSums[r] - prefixSums[l]
		}
	}
	return ans
}
