package main

func minReorder(n int, connections [][]int) int {
	var g = make([][][2]int, n) // g记录每个节点的邻接节点信息
	for _, connection := range connections {
		g[connection[0]] = append(g[connection[0]], [2]int{connection[1], 1})
		g[connection[1]] = append(g[connection[1]], [2]int{connection[0], 0})
	}
	var dfs func(int, int) int
	dfs = func(a int, fa int) (ans int) { // 参数：当前节点和当前节点的父节点，返回值：需要反转的边数
		for _, e := range g[a] {
			if e[0] != fa {
				ans += e[1] + dfs(e[0], a)
			}
		}
		return
	}
	return dfs(0, -1)
}
