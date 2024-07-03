package main

func getAncestors(n int, edges [][]int) [][]int {
	ans := make([][]int, n)
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	visit := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if visit[i] {
			return
		}
		visit[i] = true
		for _, j := range g[i] {
			if !visit[j] {
				dfs(j)
			}
		}
	}
	for i := range ans {
		clear(visit)
		dfs(i)
		visit[i] = false
		for j, b := range visit {
			if b {
				ans[i] = append(ans[i], j)
			}
		}
	}
	return ans
}
