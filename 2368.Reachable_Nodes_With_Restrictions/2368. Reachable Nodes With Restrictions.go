package main

func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	var mp = make(map[int]bool)
	for _, r := range restricted {
		mp[r] = true
	}
	var g = make([][]int, n)
	for _, edge := range edges {
		var x, y = edge[0], edge[1]
		if mp[x] || mp[y] {
			continue
		}
		g[x], g[y] = append(g[x], y), append(g[y], x)
	}
	var dfs func(x, fa int) int
	dfs = func(x, fa int) (res int) {
		for _, y := range g[x] {
			if y != fa {
				res += dfs(y, x)
			}
		}
		return res + 1
	}
	return dfs(0, -1)
}
