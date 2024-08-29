package main

func removeStones(stones [][]int) int {
	n := len(stones)
	rowMap, colMap := make(map[int][]int), make(map[int][]int)
	for i, stone := range stones {
		x, y := stone[0], stone[1]
		rowMap[x] = append(rowMap[x], i)
		colMap[y] = append(colMap[y], i)
	}
	g := make([][]int, n)
	for _, ids := range rowMap {
		for i := 1; i < len(ids); i++ {
			v, w := ids[i-1], ids[i] // 注意到任意两点间之间是直接相连还是间接相连对结果并无影响，因此考虑对于拥有 k 个石子的任意一行或一列，都恰使用 k−1 条边进行连接
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
	}
	for _, ids := range colMap {
		for i := 1; i < len(ids); i++ {
			v, w := ids[i-1], ids[i]
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
	}
	vis := make([]bool, n)
	var dfs func(v int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}
	var cnt int // 连通块数
	for i, v := range vis {
		if !v {
			cnt++
			dfs(i)
		}
	}
	return n - cnt // 对于同一个连通块，总能通过某种方式将其移除到只剩一块石头，因此总共能移除的石头总数就是石头总数减去连通块个数
}
