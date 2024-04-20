package main

func findFarmland(land [][]int) (ans [][]int) {
	var m, n = len(land), len(land[0])
	for i, row := range land {
		for j, v := range row {
			if v == 1 && (i == 0 || i > 0 && land[i-1][j] == 0) && (j == 0 || j > 0 && row[j-1] == 0) {
				var x, y = i, j
				for ; x+1 < m && land[x+1][j] == 1; x++ {
				}
				for ; y+1 < n && land[i][y+1] == 1; y++ {
				}
				ans = append(ans, []int{i, j, x, y})
			}
		}
	}
	return
}
