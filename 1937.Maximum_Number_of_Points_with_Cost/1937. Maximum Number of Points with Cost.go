package main

func maxPoints(points [][]int) int64 {
	n := len(points[0])
	f := make([][2]int, n)      // f[j][0] 和 f[j][1] 分别表示当前行第 j 个格子的最大得分值加上 j 和减去 j
	suffixMax := make([]int, n) // 后缀最大值
	var ans int
	for i, row := range points {
		if i == 0 {
			for j, point := range row {
				ans = max(ans, point)
				f[j] = [2]int{point + j, point - j}
			}
		} else {
			var prefixMax int // 前缀最大值
			for j, point := range row {
				prefixMax = max(prefixMax, f[j][0])
				res := point + max(prefixMax-j, suffixMax[j]+j) // 左侧和右侧中的最大值
				ans = max(ans, res)
				f[j] = [2]int{res + j, res - j}
			}
		}
		suffixMax[n-1] = f[n-1][1] // 在处理下一行之前先预处理本行，倒序更新后缀最大值
		for j := n - 2; j >= 0; j-- {
			suffixMax[j] = max(suffixMax[j+1], f[j][1])
		}
	}
	return int64(ans)
}
