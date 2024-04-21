package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	var g = make([][]int, n)
	for _, edge := range edges {
		var u, v = edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var visited = make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range g[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(source)
	return visited[destination]
}
