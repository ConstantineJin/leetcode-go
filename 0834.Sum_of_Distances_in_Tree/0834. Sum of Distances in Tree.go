package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var g = make([][]int, n)
	for _, edge := range edges {
		g[edge[0]], g[edge[1]] = append(g[edge[0]], edge[1]), append(g[edge[1]], edge[0])
	}
	var ans, size = make([]int, n), make([]int, n)
	var dfs func(x, fa, depth int)
	dfs = func(x, fa, depth int) {
		ans[0] += depth          // 当前节点的深度就是0号节点与当前节点的距离
		size[x] = 1              // 以0号节点为整棵树的根节点时以当前节点为根的子树的大小
		for _, y := range g[x] { // 遍历当前节点的所有邻接节点
			if y != fa { // 避免重复父节点
				dfs(y, x, depth+1)
				size[x] += size[y]
			}
		}
		return
	}
	dfs(0, -1, 0) // 以0号节点为根节点进行DFS
	var reroot func(x, fa int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				ans[y] = ans[x] + n - size[y]*2 // 在子树中的节点数有size[y]个，移动后距离都-1，不在子树中的节点数有n-size[y]个，移动后距离都+1
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
