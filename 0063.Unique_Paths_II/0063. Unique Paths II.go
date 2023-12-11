package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	path := make([][]int, m) // path记录每个格子到达右下角的路径数
	for i := range path {
		path[i] = make([]int, n)
	}
	for i := m - 1; i >= 0 && obstacleGrid[i][n-1] == 0; i-- { // 从下往上初始化最后一列，如果遇到原图中的障碍，则从这格开始往上的所有格子到达右下角的路径数均为0
		path[i][n-1] = 1
	}
	for j := n - 1; j >= 0 && obstacleGrid[m-1][j] == 0; j-- { // 从右往左初始化最后一行，如果遇到原图中的障碍，则从这格开始往左的所有格子到达右下角的路径数均为0
		path[m-1][j] = 1
	}
	for i := m - 2; i >= 0; i-- { // 从下往上，从右往左递推
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 0 { // 当前的格子不是障碍
				path[i][j] = path[i+1][j] + path[i][j+1] // 当前格子到达右下角的路径数等于当前格子右侧及下方的格子到达右下角的路径数之和
			}
		}
	}
	return path[0][0]
}
