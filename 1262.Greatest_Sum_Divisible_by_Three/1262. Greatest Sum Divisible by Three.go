package main

import "math"

func maxSumDivK(nums []int, k int) int {
	n := len(nums)
	f := make([][]int, n+1) // f[i][j] 表示 nums[:i+2] 的子序列和中 mod k = j 的最大和
	for i := range f {
		f[i] = make([]int, k)
	}
	for i := 1; i < k; i++ {
		f[0][i] = math.MinInt
	}
	for i, num := range nums {
		for j := 0; j < k; j++ {
			f[i+1][j] = max(f[i][j], f[i][(j+num)%k]+num)
		}
	}
	return f[n][0]
}

func maxSumDivThree(nums []int) int {
	return maxSumDivK(nums, 3)
}
