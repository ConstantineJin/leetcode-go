package main

import "slices"

func countGoodNodes(edges [][]int) (ans int) {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var dfs func(x, fa int) int // dfs(x,fa) 返回以 x 为根节点的子树中的节点数（x 的父节点为 fa）
	dfs = func(x, fa int) (res int) {
		sizes := make([]int, 0, len(g[x]))
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			size := dfs(y, x)
			sizes = append(sizes, size)
			res += size
		}
		if len(sizes) == 0 || slices.Min(sizes) == slices.Max(sizes) { // 叶子节点也是好节点
			ans++
		}
		return res + 1
	}
	dfs(0, -1)
	return
}
