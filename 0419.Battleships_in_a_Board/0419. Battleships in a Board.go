package main

func countBattleships(board [][]byte) (ans int) {
	for i, row := range board {
		for j, grid := range row {
			if grid == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				ans++
			}
		}
	}
	return
}
