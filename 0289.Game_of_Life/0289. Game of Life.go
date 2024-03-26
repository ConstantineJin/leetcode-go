package main

// 模拟
func gameOfLife(board [][]int) {
	var m, n = len(board), len(board[0])
	var ans = make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i, row := range board {
		for j, cell := range row {
			var live int
			if cell == 1 {
				for _, dir := range dirs {
					var x, y = i + dir[0], j + dir[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if board[x][y] == 1 {
							live++
						}
					}
				}
				if live == 2 || live == 3 {
					ans[i][j] = 1
				}
			} else {
				for _, dir := range dirs {
					var x, y = i + dir[0], j + dir[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if board[x][y] == 1 {
							live++
						}
					}
				}
				if live == 3 {
					ans[i][j] = 1
				}
			}
		}
	}
	copy(board, ans)
}
