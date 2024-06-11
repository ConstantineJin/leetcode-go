package main

// 动态规划
// f[x][j] 表示以数值 x 结尾的、有至多 j 对相邻元素不同的最长子序列的长度
// 1. 把 x 加到以 x 结尾的子序列的末尾：f[x][j] = f[x][j] + 1
// 2. 把 x 加到以 y 结尾的子序列的末尾：f[x][j] = f[y][j-1] + 1, if x != y
func maximumLength(nums []int, k int) int {
	fs := make(map[int][]int) // fs[num][j] 表示以数值 num 结尾的有至多 j 对相邻元素不同的最长子序列的长度。题目所给数据范围中 nums[i] 的分布是较稀疏的，故使用 map 而非 slice
	mx := make([]int, k+2)    // mx[j] 表示 max(f[y][j] for y in set)
	for _, num := range nums {
		if fs[num] == nil {
			fs[num] = make([]int, k+1)
		}
		f := fs[num]
		for j := k; j >= 0; j-- {
			f[j] = max(f[j], mx[j]) + 1
			mx[j+1] = max(mx[j+1], f[j])
		}
	}
	return mx[k+1]
}
