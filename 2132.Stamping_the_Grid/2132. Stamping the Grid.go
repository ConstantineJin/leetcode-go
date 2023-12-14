package main

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	var m, n = len(grid), len(grid[0])
	// 计算二维前缀和，s[i][j]表示以a[0][0]为左上角，a[i][j]为右下角的子矩阵和
	var s = make([][]int, m+1)
	s[0] = make([]int, n+1)
	for i, row := range grid {
		s[i+1] = make([]int, n+1)
		for j, v := range row {
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + v
		}
	}
	// 计算二维差分
	var d = make([][]int, m+2)
	for i := range d {
		d[i] = make([]int, n+2)
	}
	for i := stampHeight; i <= m; i++ { // i,j表示右下角坐标
		for j := stampWidth; j <= n; j++ {
			ii := i - stampHeight + 1 // ii,jj表示左上角坐标
			jj := j - stampWidth + 1
			if s[i][j]-s[i][jj-1]-s[ii-1][j]+s[ii-1][jj-1] == 0 { // 利用前缀和来计算当前子矩阵是否为全0
				d[ii][jj]++
				d[ii][j+1]--
				d[i+1][jj]--
				d[i+1][j+1]++
			}
		}
	}
	for i, row := range grid {
		for j, v := range row {
			d[i+1][j+1] += d[i+1][j] + d[i][j+1] - d[i][j]
			if v == 0 && d[i+1][j+1] == 0 { // 原来是0，盖章后仍是0，说明不能完成
				return false
			}
		}
	}
	return true
}
