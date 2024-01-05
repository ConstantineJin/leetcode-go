package main

func exist(board [][]byte, word string) (ans bool) {
	var m, n, length = len(board), len(board[0]), len(word)
	var isSearched = make([][]bool, m)
	for i := range isSearched {
		isSearched[i] = make([]bool, n)
	}
	var dfs func(x, y, l int)
	dfs = func(x, y, l int) {
		if l == length { // 找到了对应单词
			ans = true
			return
		}
		if x < 0 || y < 0 || x == m || y == n || isSearched[x][y] { // 当前搜索坐标越界或已经在搜索中
			return
		}
		isSearched[x][y] = true     // 将当前坐标标记为搜索中
		if board[x][y] == word[l] { // 如果当前坐标值与word对于字符相等，则继续对相邻坐标进行DFS
			dfs(x-1, y, l+1)
			dfs(x+1, y, l+1)
			dfs(x, y-1, l+1)
			dfs(x, y+1, l+1)
		}
		isSearched[x][y] = false // 取消对当前坐标的标记
	}
	for i := range board {
		for j := range board[i] {
			if ans { // 已经找到了单词，可以直接返回，无需继续搜索
				return
			}
			dfs(i, j, 0)
		}
	}
	return
}
