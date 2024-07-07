package main

func numberOfSubmatrices(grid [][]byte) (ans int) {
	prefixX, prefixY := make([][]int, len(grid)+1), make([][]int, len(grid)+1)
	for i := range prefixX {
		prefixX[i], prefixY[i] = make([]int, len(grid[0])+1), make([]int, len(grid[0])+1)
	}
	for i, row := range grid {
		for j, v := range row {
			prefixX[i+1][j+1] = prefixX[i+1][j] + prefixX[i][j+1] - prefixX[i][j]
			prefixY[i+1][j+1] = prefixY[i+1][j] + prefixY[i][j+1] - prefixY[i][j]
			switch v {
			case 'X':
				prefixX[i+1][j+1]++
			case 'Y':
				prefixY[i+1][j+1]++
			}
			if prefixX[i+1][j+1] > 0 && prefixX[i+1][j+1] == prefixY[i+1][j+1] {
				ans++
			}
		}
	}
	return
}
