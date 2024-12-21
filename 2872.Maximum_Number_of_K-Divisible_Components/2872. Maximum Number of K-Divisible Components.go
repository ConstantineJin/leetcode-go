package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) (ans int) {
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	var dfs func(x, fa int) int
	dfs = func(x, fa int) int {
		s := values[x]
		for _, y := range g[x] {
			if y != fa {
				s += dfs(y, x)
			}
		}
		if s%k == 0 {
			ans++
		}
		return s
	}
	dfs(0, -1)
	return
}
