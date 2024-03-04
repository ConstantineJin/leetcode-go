package main

func solve(board [][]byte) {
	var m, n = len(board), len(board[0])
	var ans = make([][]byte, m)
	for i := range ans {
		ans[i] = make([]byte, n)
		for j := range ans[i] {
			ans[i][j] = 'X'
		}
	}
	var start [][2]int
	for i := range board {
		if board[i][0] == 'O' {
			start = append(start, [2]int{i, 0})
		}
		if board[i][n-1] == 'O' {
			start = append(start, [2]int{i, n - 1})
		}
	}
	for i := range board[0] {
		if i > 0 && i < n-1 && board[0][i] == 'O' {
			start = append(start, [2]int{0, i})
		}
		if i > 0 && i < n-1 && board[m-1][i] == 'O' {
			start = append(start, [2]int{m - 1, i})
		}
	}
	var dirs = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var seen = make(map[[2]int]bool)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || seen[[2]int{i, j}] {
			return
		}
		seen[[2]int{i, j}] = true
		if board[i][j] == 'O' {
			ans[i][j] = 'O'
			for _, dir := range dirs {
				dfs(i+dir[0], j+dir[1])
			}
		}
	}
	for _, s := range start {
		clear(seen)
		dfs(s[0], s[1])
	}
	copy(board, ans)
}
