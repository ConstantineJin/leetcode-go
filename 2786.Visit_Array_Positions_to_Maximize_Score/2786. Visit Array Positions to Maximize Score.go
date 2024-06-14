package main

// 动态规划，递推
//func maxScore(nums []int, x int) int64 {
//	n := len(nums)
//	f := make([][2]int, n+1) // f[i][j] 表示在下标 [i,n−1] 中选一个子序列，其第一个数的奇偶性为 j（即模 2 的结果为 j）时的最大得分
//	for i := n - 1; i >= 0; i-- {
//		r := nums[i] % 2
//		f[i][r^1] = f[i+1][r^1] // nums[i]%2 != j 的情况
//		f[i][r] = max(f[i+1][r], f[i+1][r^1]-x) + nums[i]
//	}
//	return int64(f[0][nums[0]%2])
//}

// 空间复杂度优化至O(1)
func maxScore(nums []int, x int) int64 {
	var f [2]int
	for i := len(nums) - 1; i >= 0; i-- {
		r := nums[i] % 2
		f[r] = max(f[r], f[r^1]-x) + nums[i]
	}
	return int64(f[nums[0]%2])
}
