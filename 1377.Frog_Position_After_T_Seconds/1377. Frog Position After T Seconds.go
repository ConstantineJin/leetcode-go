package main

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1) // 邻接表
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	visited := make([]bool, n+1)
	visited[0], visited[1] = true, true
	probabilities := make([]float64, n+1)
	probabilities[1] = 1
	cur := []int{1} // BFS模拟
	for time := 0; time < t; time++ {
		var nxt []int
		for _, x := range cur {
			var cnt int // 当前节点的邻接节点中未访问过的节点数
			for _, y := range g[x] {
				if !visited[y] {
					cnt++
					nxt = append(nxt, y)
				}
			}
			if cnt != 0 {
				for _, y := range g[x] {
					if !visited[y] {
						probabilities[y] = probabilities[x] / float64(cnt) // 每个未访问的邻接节点均分当前节点的概率
						visited[y] = true
					}
				}
				probabilities[x] = 0 // 当前节点的概率变为0
			} // 如果没有未访问的邻接节点，那么当前节点的概率从此以后不会发生任何变化
		}
		cur = nxt
	}
	return probabilities[target]
}
