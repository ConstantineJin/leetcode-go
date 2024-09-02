package main

func maximumSubarrayXor(nums []int, queries [][]int) []int {
	n := len(nums)
	f := make([][]int, n)  // f[i][j] 表示子数组 nums[i:j] 的异或值
	mx := make([][]int, n) // mx[i][j] 表示 nums[i:j] 的所有子数组的最大异或值
	for i := range f {
		f[i], mx[i] = make([]int, n), make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i], mx[i][i] = nums[i], nums[i]
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i][j-1] ^ f[i+1][j]
			mx[i][j] = max(f[i][j], mx[i+1][j], mx[i][j-1])
		}
	}
	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		ans = append(ans, mx[query[0]][query[1]])
	}
	return ans
}
