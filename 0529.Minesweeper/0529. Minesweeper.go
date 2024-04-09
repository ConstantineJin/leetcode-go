package main

var dirs = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	var m, n, x, y = len(board), len(board[0]), click[0], click[1]
	switch board[x][y] {
	case 'M':
		board[x][y] = 'X'
	case 'E':
		var cnt int
		for _, dir := range dirs {
			var xx, yy = x + dir[0], y + dir[1]
			if xx > -1 && xx < m && yy > -1 && yy < n && board[xx][yy] == 'M' {
				cnt++
			}
		}
		if cnt == 0 {
			board[x][y] = 'B'
			for _, dir := range dirs {
				var xx, yy = x + dir[0], y + dir[1]
				if xx > -1 && xx < m && yy > -1 && yy < n {
					updateBoard(board, []int{xx, yy})
				}
			}
		} else {
			board[x][y] = byte('0' + cnt)
		}
	}
	return board
}
