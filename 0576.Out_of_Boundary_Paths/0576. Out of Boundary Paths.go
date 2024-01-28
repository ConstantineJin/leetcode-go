package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findPaths(m, n, maxMove, startRow, startColumn int) (ans int) {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	for i := 0; i < maxMove; i++ {
		dpNew := make([][]int, m)
		for j := range dpNew {
			dpNew[j] = make([]int, n)
		}
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				count := dp[j][k]
				if count > 0 {
					for _, dir := range dirs { // 分别向上下左右四个方向移动
						j1, k1 := j+dir.x, k+dir.y                  // 新坐标
						if j1 >= 0 && j1 < m && k1 >= 0 && k1 < n { // 未出界
							dpNew[j1][k1] = (dpNew[j1][k1] + count) % (1e9 + 7) // 到达[j1,k1]的路径数
						} else { // 出界
							ans = (ans + count) % (1e9 + 7) // 出界路径数添加到结果
						}
					}
				}
			}
		}
		dp = dpNew // 空间复杂度降维，O(n^3)->O(n^2)
	}
	return
}
