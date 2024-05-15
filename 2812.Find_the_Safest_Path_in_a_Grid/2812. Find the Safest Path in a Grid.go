package main

type cell struct{ x, y int }

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	distance := make([][]int, n) // 每个单元格到最近的小偷的曼哈顿距离
	var cur []cell
	for i, row := range grid {
		distance[i] = make([]int, n)
		for j, v := range row {
			if v == 1 {
				cur = append(cur, cell{i, j})
			} else {
				distance[i][j] = -1
			}
		}
	}
	groups := [][]cell{cur}
	// 多源BFS
	for dis := 1; len(cur) > 0; dis++ {
		var nxt []cell
		for _, c := range cur {
			for _, dir := range dirs {
				if x, y := c.x+dir[0], c.y+dir[1]; x >= 0 && x < n && y >= 0 && y < n && distance[x][y] == -1 {
					distance[x][y] = dis
					nxt = append(nxt, cell{x, y})
				}
			}
		}
		cur = nxt
		groups = append(groups, nxt)
	}
	// 并查集
	fa := make([]int, n*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for ans := len(groups) - 2; ans > 0; ans-- { // 假设答案为d，将所有distance[i][j]==d的格子与其四周所有>=d的格子使用并查集相连
		for _, c := range groups[ans] {
			for _, dir := range dirs {
				if x, y := c.x+dir[0], c.y+dir[1]; x >= 0 && x < n && y >= 0 && y < n && distance[x][y] >= distance[c.x][c.y] {
					fa[find(x*n+y)] = find(c.x*n + c.y)
				}
			}
		}
		if find(0) == find(n*n-1) { // 一旦左上角和右下角联通就立刻返回
			return ans
		}
	}
	return 0
}
