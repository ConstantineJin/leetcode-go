package main

func timeTaken(edges [][]int) []int {
	g := make([][]int, len(edges)+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	nodes := make([]struct{ maxD, maxD2, y int }, len(g)) // nodes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 y 走
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		p := &nodes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			maxD := dfs(y, x) + 2 - y%2 // 从 x 出发，往 y 方向的最大深度
			if maxD > p.maxD {
				p.maxD2 = p.maxD
				p.maxD = maxD
				p.y = y
			} else if maxD > p.maxD2 {
				p.maxD2 = maxD
			}
		}
		return p.maxD
	}
	dfs(0, -1)

	ans := make([]int, len(g))
	var reRoot func(int, int, int) // 换根DP
	reRoot = func(x, fa, fromUp int) {
		p := nodes[x]
		ans[x] = max(fromUp, p.maxD)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			w := 2 - x%2  // 从 y 到 x 的边权
			if y == p.y { // 对于 y 来说，上面要选次大的
				reRoot(y, x, max(fromUp, p.maxD2)+w)
			} else { // 对于 y 来说，上面要选最大的
				reRoot(y, x, max(fromUp, p.maxD)+w)
			}
		}
	}
	reRoot(0, -1, 0)
	return ans
}
