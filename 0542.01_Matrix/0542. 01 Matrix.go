package main

type pair struct{ x, y int }

var dirs = [4]pair{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}
	var cur []pair
	for i, row := range mat {
		for j, v := range row {
			if v == 0 {
				vis[i][j] = true
				cur = append(cur, pair{i, j})
			}
		}
	}
	for len(cur) > 0 {
		var nxt []pair
		for _, g := range cur {
			for _, dir := range dirs {
				if x, y := g.x+dir.x, g.y+dir.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] {
					mat[x][y] = mat[g.x][g.y] + 1
					vis[x][y] = true
					nxt = append(nxt, pair{x, y})
				}
			}
		}
		cur = nxt
	}
	return mat
}
