package main

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	type pair struct{ id, weight int }
	g := make([][]pair, n) // 建图
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], pair{edge[1], edge[2]})
		g[edge[1]] = append(g[edge[1]], pair{edge[0], edge[2]})
	}
	ans := make([]int, n)
	for i := range ans { // 遍历每个节点作为根节点的情况
		arr := make([]int, len(g[i])) // 记录当前根节点的每个子树中可以和根节点连接的服务器数量
		for j, node := range g[i] {
			var dfs func(id, fa, dis int) // 深度优先搜索，id为当前搜索到的节点编号，fa为当前节点的父节点，dis为从根节点出发到当前节点的距离
			dfs = func(id, fa, dis int) {
				if dis%signalSpeed == 0 {
					arr[j]++
				}
				for _, node := range g[id] {
					if node.id != fa {
						dfs(node.id, id, dis+node.weight)
					}
				}
			}
			dfs(node.id, i, node.weight)
		}
		// 乘法原理
		for j := 0; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				ans[i] += arr[j] * arr[k]
			}
		}
	}
	return ans
}
