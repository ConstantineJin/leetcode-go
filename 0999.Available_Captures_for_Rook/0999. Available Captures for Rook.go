package main

var dirs = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func numRookCaptures(board [][]byte) (ans int) {
	var x, y int
out:
	for i, row := range board {
		for j, c := range row {
			if c == 'R' {
				x, y = i, j
				break out
			}
		}
	}
next:
	for _, dir := range dirs {
		for i, j := x+dir[0], y+dir[1]; 0 <= i && i < 8 && 0 <= j && j < 8 && board[i][j] != 'B'; i, j = i+dir[0], j+dir[1] {
			if board[i][j] == 'p' {
				ans++
				continue next
			}
		}
	}
	return
}
