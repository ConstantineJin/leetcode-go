package main

import "slices"

func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n) // great[k][x] 表示在 k 右侧的比 x=nums[j] 大的元素个数
	great[n-1] = make([]int, n+1)
	for k := n - 2; k > 0; k-- {
		great[k] = slices.Clone(great[k+1])
		for x := 1; x < nums[k+1]; x++ {
			great[k][x]++
		}
	}
	for j := 1; j < n-2; j++ { // 枚举中间两个元素
		for k := j + 1; k < n-1; k++ {
			x := nums[k]
			if nums[j] > x {
				ans += int64((x - n + 1 + j + great[j][x]) * great[k][nums[j]]) // 乘法原理。由于 nums 是一个 [1,n] 的排列，因此 j 右边有 n-1-j-great[j][x] 个不超过 x 的数。又因为总共有 x 个不超过 x 的数，且 x 在 j 的右侧，那么 j 左侧小于 x 的个数为 x-(n-1-j-great[j][x])
			}
		}
	}
	return
}
