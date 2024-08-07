package main

import "sort"

const mod = 1e9 + 7

func rangeSum(nums []int, n, left, right int) (ans int) {
	prefixSum := make([]int, 0)
	var curSum int
	for i := 0; i <= n; i++ {
		prefixSum = append(prefixSum, curSum)
		if i == n {
			break
		}
		curSum += nums[i]
	}
	var allSums []int
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			allSums = append(allSums, prefixSum[j]-prefixSum[i])
		}
	}
	sort.Ints(allSums)
	for i := left - 1; i <= right-1; i++ {
		ans = (ans + allSums[i]) % mod
	}
	return ans
}
