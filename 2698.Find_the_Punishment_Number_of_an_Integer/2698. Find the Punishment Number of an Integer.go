package main

import "sort"

func punishmentNumber(n int) (ans int) {
	var arr = []int{1, 9, 10, 36, 45, 55, 82, 91, 99, 100, 235, 297, 369, 370, 379, 414, 657, 675, 703, 756, 792, 909, 918, 945, 964, 990, 991, 999, 1000}
	idx := sort.SearchInts(arr, n)
	if arr[idx] != n {
		idx--
	}
	for i := 0; i <= idx; i++ {
		ans += arr[i] * arr[i]
	}
	return ans
}
