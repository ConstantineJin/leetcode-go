package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	g := make([][]int, n)
	for i, bomb0 := range bombs {
		x, y, r := bomb0[0], bomb0[1], bomb0[2]
		for j, bomb1 := range bombs {
			dx, dy := abs(x-bomb1[0]), abs(y-bomb1[1])
			if i != j && dx*dx+dy*dy <= r*r {
				g[i] = append(g[i], j)
			}
		}
	}
	visited := make([]bool, n)
	var dfs func(i int) (res int)
	dfs = func(i int) (res int) {
		visited[i] = true
		res = 1
		for _, j := range g[i] {
			if !visited[j] {
				res += dfs(j)
			}
		}
		return
	}
	for i := range g {
		clear(visited)
		ans = max(ans, dfs(i))
	}
	return
}
