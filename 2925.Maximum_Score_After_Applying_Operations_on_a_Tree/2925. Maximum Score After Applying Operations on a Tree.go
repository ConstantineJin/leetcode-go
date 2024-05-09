package main

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	g := make([][]int, len(values))
	g[0] = append(g[0], -1) // 避免误把根节点当作叶子
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var ans int
	for _, v := range values {
		ans += v
	}
	// dfs(x, fa) 计算以 x 为根的子树是健康时，失去的最小分数
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		if len(g[x]) == 1 { // x 是叶子
			return values[x]
		}
		var loss int // 不选 values[x]
		for _, y := range g[x] {
			if y != fa {
				loss += dfs(y, x) // 计算以 y 为根的子树是健康时，失去的最小分数
			}
		}
		return min(values[x], loss) // 选/不选 values[x]，取最小值
	}
	return int64(ans - dfs(0, -1))
}
