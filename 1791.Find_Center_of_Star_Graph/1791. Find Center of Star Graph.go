package main

func findCenter(edges [][]int) int {
	g := make([]int, len(edges)+2) // 星型图有 n 个节点，则有 n-1 条边。本题下标从 1 开始
	for _, edge := range edges {
		g[edge[0]]++
		if g[edge[0]] > 1 {
			return edge[0]
		}
		g[edge[1]]++
		if g[edge[1]] > 1 {
			return edge[1]
		}
	}
	return 0
}
