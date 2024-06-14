package main

func solveSudoku(board [][]byte) {
	var rows, cols [9][9]bool // rows[i][j]/cols[i][j] 表示第 i 行/列已有数字 j
	var block [3][3][9]bool   // block[i][j][k] 表示从上往下第 i 个从左往右第 j 个九宫格中已有数字 k
	var spaces [][2]int       // 数独中初始为空的单元格
	for i, row := range board {
		for j, grid := range row {
			if grid == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				rows[i][grid-'1'] = true
				cols[j][grid-'1'] = true
				block[i/3][j/3][grid-'1'] = true
			}
		}
	}
	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for v := 0; v < 9; v++ {
			if !rows[i][v] && !cols[j][v] && !block[i/3][j/3][v] {
				rows[i][v], cols[j][v], block[i/3][j/3][v] = true, true, true
				board[i][j] = byte(v + '1')
				if dfs(pos + 1) {
					return true
				}
				rows[i][v], cols[j][v], block[i/3][j/3][v] = false, false, false
			}
		}
		return false
	}
	dfs(0)
}
