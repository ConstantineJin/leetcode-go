package main

// 方法1:二维前缀和
func numSubmat(mat [][]int) (ans int) {
	m, n := len(mat), len(mat[0])
	prefixSums := make([][]int, m+1)
	for i := range prefixSums {
		prefixSums[i] = make([]int, n+1)
	}
	for i, row := range mat {
		for j, v := range row {
			prefixSums[i+1][j+1] = prefixSums[i+1][j] + prefixSums[i][j+1] - prefixSums[i][j] + v
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for x := i - 1; x >= 0; x-- {
				for y := j - 1; y >= 0; y-- {
					if prefixSums[i][j]-prefixSums[x][j]-prefixSums[i][y]+prefixSums[x][y] < (i-x)*(j-y) {
						break
					}
					ans++
				}
			}
		}
	}
	return
}
