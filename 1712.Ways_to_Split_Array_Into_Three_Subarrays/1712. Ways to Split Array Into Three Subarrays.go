package main

import "sort"

func waysToSplit(nums []int) (ans int) {
	n := len(nums)
	prefixSums := make([]int, n+1)
	for i, num := range nums {
		prefixSums[i+1] = prefixSums[i] + num
	}
	for r := 2; r < n && 3*prefixSums[r] <= 2*prefixSums[n]; r++ {
		l := sort.SearchInts(prefixSums[1:r], 2*prefixSums[r]-prefixSums[n]) + 1 // S[r]-S[l]≤S[n]-S[r]
		ans += sort.SearchInts(prefixSums[l:r], prefixSums[r]/2+1)               // S[l]≤S[r]-S[l]
	}
	return ans % (1e9 + 7)
}
