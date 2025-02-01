package main

var dirs = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func largestIsland(grid [][]int) (ans int) {
	n := len(grid)
	var areas []int
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		grid[i][j] = len(areas) + 2 // 记录 (i,j) 属于哪个岛
		size := 1
		for _, dir := range dirs {
			if x, y := i+dir[0], j+dir[1]; 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
				size += dfs(x, y)
			}
		}
		return size
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				areas = append(areas, dfs(i, j))
			}
		}
	}
	if len(areas) == 0 {
		return 1
	}
	s := make(map[int]bool)
	for i, row := range grid {
		for j, v := range row {
			if v != 0 {
				continue
			}
			clear(s)
			newArea := 1
			for _, dir := range dirs {
				if x, y := i+dir[0], j+dir[1]; 0 <= x && x < n && 0 <= y && y < n && grid[x][y] != 0 && !s[grid[x][y]] {
					s[grid[x][y]] = true
					newArea += areas[grid[x][y]-2]
				}
			}
			ans = max(ans, newArea)
		}
	}
	if ans == 0 {
		return n * n
	}
	return
}
