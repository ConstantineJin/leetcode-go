package main

func getDiameter(edges [][]int) (res int) {
	g := make([][]int, len(edges)+1)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	var dfs func(x, fa int) (maxLen int)
	dfs = func(x, fa int) (maxLen int) {
		for _, y := range g[x] {
			if y != fa {
				subLen := dfs(y, x) + 1
				res = max(res, maxLen+subLen)
				maxLen = max(maxLen, subLen)
			}
		}
		return
	}
	dfs(0, -1)
	return
}

func minimumDiameterAfterMerge(edges1, edges2 [][]int) int {
	d1, d2 := getDiameter(edges1), getDiameter(edges2)
	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
}
