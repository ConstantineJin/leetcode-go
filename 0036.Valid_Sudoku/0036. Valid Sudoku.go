package main

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subBoxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			var index = c - '1'
			rows[i][index]++
			columns[j][index]++
			subBoxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subBoxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
