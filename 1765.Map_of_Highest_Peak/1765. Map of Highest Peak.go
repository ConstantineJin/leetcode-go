package main

type pair struct{ x, y int }

var dirs = [4]pair{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
		for j := range n {
			ans[i][j] = -1
		}
	}
	var cur []pair
	for i, row := range isWater {
		for j, v := range row {
			if v == 1 {
				ans[i][j] = 0
				cur = append(cur, pair{i, j})
			}
		}
	}
	for len(cur) > 0 {
		var nxt []pair
		for _, g := range cur {
			for _, dir := range dirs {
				if x, y := g.x+dir.x, g.y+dir.y; 0 <= x && x < m && 0 <= y && y < n && ans[x][y] == -1 {
					ans[x][y] = ans[g.x][g.y] + 1
					nxt = append(nxt, pair{x, y})
				}
			}
		}
		cur = nxt
	}
	return ans
}
